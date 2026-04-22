//ff:func feature=migration type=parser control=sequence
//ff:what BuildAST — *.sql 파일들의 원본 텍스트를 Schema AST 로 파싱 (기존 DDL 파서에 의존 X)
package migration

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// BuildASTFromDir walks <dir> for *.sql files and returns a Schema.
// It intentionally re-reads the raw files rather than going through
// pkg/parser/ddl because that parser already lossily converts column
// types to Go types. For diff purposes we need the full PostgreSQL type
// information (INT vs BIGINT vs VARCHAR(255) …).
//
// skipFiles lists base file names to skip (e.g. ".generated_schema.sql").
func BuildASTFromDir(dir string, skipFiles []string) (*Schema, error) {
	skip := map[string]bool{}
	for _, f := range skipFiles {
		skip[f] = true
	}
	s := NewSchema()
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return s, nil
		}
		return nil, fmt.Errorf("read dir %s: %w", dir, err)
	}
	var files []string
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		name := e.Name()
		if !strings.HasSuffix(strings.ToLower(name), ".sql") {
			continue
		}
		if skip[name] {
			continue
		}
		files = append(files, filepath.Join(dir, name))
	}
	sort.Strings(files)
	for _, f := range files {
		data, err := os.ReadFile(f)
		if err != nil {
			return nil, fmt.Errorf("read %s: %w", f, err)
		}
		if err := BuildASTFromSQL(s, string(data)); err != nil {
			return nil, fmt.Errorf("parse %s: %w", f, err)
		}
	}
	return s, nil
}

// BuildASTFromSQL appends the structures parsed from one SQL text into
// the given Schema. It is the main lexer/parser entry point and is
// exported so tests can feed SQL literals directly.
func BuildASTFromSQL(s *Schema, sqlText string) error {
	// Remove -- line comments but keep them for @archived detection when
	// needed (v1 migration does not consume @archived — separate path).
	stmts := splitStatements(sqlText)
	for _, stmt := range stmts {
		stmt = strings.TrimSpace(stmt)
		if stmt == "" {
			continue
		}
		upper := strings.ToUpper(stmt)
		switch {
		case strings.HasPrefix(upper, "CREATE TABLE"):
			if err := parseCreateTable(s, stmt); err != nil {
				return err
			}
		case strings.HasPrefix(upper, "CREATE UNIQUE INDEX"),
			strings.HasPrefix(upper, "CREATE INDEX"):
			if err := parseCreateIndex(s, stmt); err != nil {
				return err
			}
		case strings.HasPrefix(upper, "ALTER TABLE"):
			if err := parseAlterTable(s, stmt); err != nil {
				return err
			}
			// INSERT / COMMENT / SET are ignored for schema purposes.
		}
	}
	return nil
}

// splitStatements divides SQL text on top-level `;` boundaries while
// respecting single-quoted strings, double-quoted identifiers and
// block comments. Line (`--`) comments are stripped before splitting.
func splitStatements(sql string) []string {
	clean := stripLineComments(sql)
	out := []string{}
	var sb strings.Builder
	inSQ, inDQ, inBC := false, false, false
	depth := 0
	for i := 0; i < len(clean); i++ {
		c := clean[i]
		switch {
		case inBC:
			sb.WriteByte(c)
			if c == '*' && i+1 < len(clean) && clean[i+1] == '/' {
				sb.WriteByte('/')
				i++
				inBC = false
			}
		case inSQ:
			sb.WriteByte(c)
			if c == '\'' {
				if i+1 < len(clean) && clean[i+1] == '\'' {
					sb.WriteByte('\'')
					i++
				} else {
					inSQ = false
				}
			}
		case inDQ:
			sb.WriteByte(c)
			if c == '"' {
				inDQ = false
			}
		case c == '\'':
			inSQ = true
			sb.WriteByte(c)
		case c == '"':
			inDQ = true
			sb.WriteByte(c)
		case c == '/' && i+1 < len(clean) && clean[i+1] == '*':
			inBC = true
			sb.WriteByte(c)
			sb.WriteByte('*')
			i++
		case c == '(':
			depth++
			sb.WriteByte(c)
		case c == ')':
			if depth > 0 {
				depth--
			}
			sb.WriteByte(c)
		case c == ';' && depth == 0:
			out = append(out, sb.String())
			sb.Reset()
		default:
			sb.WriteByte(c)
		}
	}
	if strings.TrimSpace(sb.String()) != "" {
		out = append(out, sb.String())
	}
	return out
}

func stripLineComments(sql string) string {
	var sb strings.Builder
	lines := strings.Split(sql, "\n")
	for _, ln := range lines {
		idx := findLineCommentStart(ln)
		if idx >= 0 {
			sb.WriteString(ln[:idx])
		} else {
			sb.WriteString(ln)
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

// findLineCommentStart returns the byte index of the `--` that starts
// an SQL line comment, or -1 if the line has no such comment. Respects
// single-quoted string literals.
func findLineCommentStart(line string) int {
	inSQ := false
	for i := 0; i < len(line)-1; i++ {
		c := line[i]
		switch c {
		case '\'':
			if inSQ && i+1 < len(line) && line[i+1] == '\'' {
				i++
				continue
			}
			inSQ = !inSQ
		case '-':
			if !inSQ && line[i+1] == '-' {
				return i
			}
		}
	}
	return -1
}

var reCreateTable = regexp.MustCompile(`(?is)^\s*CREATE\s+TABLE\s+(?:IF\s+NOT\s+EXISTS\s+)?([^\s(]+)\s*\((.*)\)\s*$`)

func parseCreateTable(s *Schema, stmt string) error {
	m := reCreateTable.FindStringSubmatch(stmt)
	if m == nil {
		return fmt.Errorf("unparseable CREATE TABLE statement: %q", trimForErr(stmt))
	}
	rawName := m[1]
	body := m[2]
	name := canonIdent(rawName)
	t, exists := s.Tables[name]
	if !exists {
		t = &Table{Name: name}
		s.Tables[name] = t
	}

	items := splitTopLevel(body, ',')
	for _, it := range items {
		it = strings.TrimSpace(it)
		if it == "" {
			continue
		}
		upper := strings.ToUpper(it)
		switch {
		case strings.HasPrefix(upper, "PRIMARY KEY"):
			t.PrimaryKey = parseColumnList(afterKeyword(it, "PRIMARY KEY"))
		case strings.HasPrefix(upper, "UNIQUE"):
			cols := parseColumnList(afterKeyword(it, "UNIQUE"))
			t.Indexes = append(t.Indexes, &Index{
				Name: UniqueName(t.Name, cols), Columns: cols, Unique: true,
			})
		case strings.HasPrefix(upper, "FOREIGN KEY"):
			fk := parseTableFK(t, it)
			if fk != nil {
				t.ForeignKeys = append(t.ForeignKeys, fk)
			}
		case strings.HasPrefix(upper, "CHECK"):
			t.Checks = append(t.Checks, parseTableCheck(t, "", it))
		case strings.HasPrefix(upper, "CONSTRAINT"):
			parseNamedConstraint(t, it)
		default:
			parseColumn(t, it)
		}
	}
	return nil
}

func afterKeyword(s, kw string) string {
	idx := strings.Index(strings.ToUpper(s), strings.ToUpper(kw))
	if idx < 0 {
		return s
	}
	return strings.TrimSpace(s[idx+len(kw):])
}

// parseColumn handles a single column definition item inside CREATE TABLE(...).
func parseColumn(t *Table, def string) {
	// Tokenize while keeping quoted identifiers and parenthesised groups
	// intact.
	tokens := tokenizeColumnDef(def)
	if len(tokens) < 2 {
		return
	}
	name := canonIdent(tokens[0])
	typeTok, rest := collectTypeTokens(tokens[1:])
	ct, isSerial := NormalizeType(typeTok)

	col := &Column{
		Name:     name,
		Type:     ct,
		Nullable: true,
	}

	// Walk the remaining tokens for NOT NULL / NULL / DEFAULT / PRIMARY
	// KEY / UNIQUE / REFERENCES / CHECK.
	i := 0
	for i < len(rest) {
		tok := strings.ToUpper(rest[i])
		switch {
		case tok == "NOT" && i+1 < len(rest) && strings.ToUpper(rest[i+1]) == "NULL":
			col.Nullable = false
			i += 2
		case tok == "NULL":
			col.Nullable = true
			i++
		case tok == "DEFAULT":
			def, consumed := collectDefaultExpr(rest[i+1:])
			col.Default = NormalizeDefault(def)
			i += 1 + consumed
		case tok == "PRIMARY" && i+1 < len(rest) && strings.ToUpper(rest[i+1]) == "KEY":
			t.PrimaryKey = []string{col.Name}
			col.Nullable = false
			i += 2
		case tok == "UNIQUE":
			t.Indexes = append(t.Indexes, &Index{
				Name:    UniqueName(t.Name, []string{col.Name}),
				Columns: []string{col.Name},
				Unique:  true,
			})
			i++
		case tok == "REFERENCES":
			fk, consumed := parseInlineRef(t, col.Name, rest[i+1:])
			if fk != nil {
				t.ForeignKeys = append(t.ForeignKeys, fk)
			}
			i += 1 + consumed
		case tok == "CHECK":
			// Inline CHECK (expr)
			if i+1 < len(rest) && strings.HasPrefix(rest[i+1], "(") {
				t.Checks = append(t.Checks, &CheckConstraint{
					Name:       CheckName(t.Name, col.Name),
					Expression: innerParens(rest[i+1]),
				})
				i += 2
			} else {
				i++
			}
		default:
			i++
		}
	}

	// SERIAL variants: attach nextval() default (canonical form uses
	// anonymous sequence name to stay comparable). We don't know the
	// sequence name deterministically, so leave default empty unless
	// user already had one.
	if isSerial && col.Default == "" {
		col.Default = "nextval('" + t.Name + "_" + col.Name + "_seq')"
		col.Nullable = false
	}

	t.Columns = append(t.Columns, col)
}

// tokenizeColumnDef splits a column definition into whitespace-separated
// tokens while keeping parenthesised groups and quoted strings as a
// single token each.
func tokenizeColumnDef(s string) []string {
	var out []string
	var sb strings.Builder
	depth := 0
	inSQ, inDQ := false, false
	flush := func() {
		if sb.Len() > 0 {
			out = append(out, sb.String())
			sb.Reset()
		}
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case inSQ:
			sb.WriteByte(c)
			if c == '\'' {
				if i+1 < len(s) && s[i+1] == '\'' {
					sb.WriteByte('\'')
					i++
				} else {
					inSQ = false
				}
			}
		case inDQ:
			sb.WriteByte(c)
			if c == '"' {
				inDQ = false
			}
		case c == '\'':
			inSQ = true
			sb.WriteByte(c)
		case c == '"':
			inDQ = true
			sb.WriteByte(c)
		case c == '(':
			depth++
			sb.WriteByte(c)
		case c == ')':
			depth--
			sb.WriteByte(c)
		case (c == ' ' || c == '\t' || c == '\n' || c == '\r') && depth == 0:
			flush()
		default:
			sb.WriteByte(c)
		}
	}
	flush()
	return out
}

// collectTypeTokens consumes one or more tokens forming a SQL type
// expression (e.g. "character varying(255)", "timestamp with time
// zone", "varchar(255)", "int[]"). Returns the joined raw type string
// and the remaining tokens.
func collectTypeTokens(toks []string) (string, []string) {
	if len(toks) == 0 {
		return "", nil
	}
	typeParts := []string{toks[0]}
	i := 1
	// Multi-word known types.
	upper := strings.ToUpper(toks[0])
	if (upper == "CHARACTER" || upper == "TIMESTAMP" || upper == "TIME" || upper == "DOUBLE") && i < len(toks) {
		next := strings.ToUpper(toks[i])
		if next == "VARYING" || next == "PRECISION" {
			typeParts = append(typeParts, toks[i])
			i++
		} else if next == "WITH" || next == "WITHOUT" {
			// "WITH TIME ZONE" / "WITHOUT TIME ZONE"
			if i+2 < len(toks) &&
				strings.ToUpper(toks[i+1]) == "TIME" &&
				strings.ToUpper(toks[i+2]) == "ZONE" {
				typeParts = append(typeParts, toks[i], toks[i+1], toks[i+2])
				i += 3
			}
		}
	}
	// Array `[]` suffix may be a separate token.
	for i < len(toks) && toks[i] == "[]" {
		typeParts[len(typeParts)-1] += "[]"
		i++
	}
	return strings.Join(typeParts, " "), toks[i:]
}

// collectDefaultExpr consumes tokens forming a DEFAULT expression.
// It stops at the next known column-level keyword (NOT, NULL, UNIQUE,
// PRIMARY, REFERENCES, CHECK, DEFAULT) and returns the joined text plus
// the number of tokens consumed.
func collectDefaultExpr(toks []string) (string, int) {
	stop := map[string]bool{
		"NOT": true, "NULL": true, "UNIQUE": true, "PRIMARY": true,
		"REFERENCES": true, "CHECK": true, "DEFAULT": true, "CONSTRAINT": true,
	}
	var parts []string
	i := 0
	for ; i < len(toks); i++ {
		if stop[strings.ToUpper(toks[i])] {
			break
		}
		parts = append(parts, toks[i])
	}
	return strings.Join(parts, " "), i
}

func parseInlineRef(t *Table, col string, toks []string) (*ForeignKey, int) {
	// REFERENCES <table>[(<col>)] [ON DELETE X] [ON UPDATE Y]
	if len(toks) == 0 {
		return nil, 0
	}
	target := toks[0]
	refTable := canonIdent(target)
	var refCol string
	consumed := 1
	if p := strings.Index(target, "("); p >= 0 {
		// e.g. "organizations(id)"
		refTable = canonIdent(target[:p])
		if end := strings.LastIndex(target, ")"); end > p {
			refCol = canonIdent(strings.TrimSpace(target[p+1 : end]))
		}
	} else if len(toks) > 1 && strings.HasPrefix(toks[1], "(") {
		refCol = canonIdent(innerParens(toks[1]))
		consumed++
	}
	fk := &ForeignKey{
		Name:       FKName(t.Name, []string{col}),
		Columns:    []string{col},
		RefTable:   refTable,
		RefColumns: []string{refCol},
	}
	// ON DELETE / ON UPDATE
	for consumed+2 < len(toks) {
		if strings.ToUpper(toks[consumed]) != "ON" {
			break
		}
		action := strings.ToUpper(toks[consumed+1])
		val, step := collectOnAction(toks[consumed+2:])
		switch action {
		case "DELETE":
			fk.OnDelete = val
		case "UPDATE":
			fk.OnUpdate = val
		default:
			break
		}
		consumed += 2 + step
	}
	return fk, consumed
}

func collectOnAction(toks []string) (string, int) {
	if len(toks) == 0 {
		return "", 0
	}
	upper := strings.ToUpper(toks[0])
	if upper == "SET" && len(toks) > 1 {
		return "SET " + strings.ToUpper(toks[1]), 2
	}
	if upper == "NO" && len(toks) > 1 && strings.ToUpper(toks[1]) == "ACTION" {
		return "NO ACTION", 2
	}
	return strings.ToUpper(toks[0]), 1
}

func parseTableFK(t *Table, item string) *ForeignKey {
	// FOREIGN KEY (a,b) REFERENCES other(c,d) [ON ...]
	rest := afterKeyword(item, "FOREIGN KEY")
	toks := tokenizeColumnDef(rest)
	if len(toks) < 3 {
		return nil
	}
	localCols := parseColumnList(toks[0])
	if strings.ToUpper(toks[1]) != "REFERENCES" {
		return nil
	}
	target := toks[2]
	refTable := canonIdent(target)
	var refCols []string
	consumed := 3
	if p := strings.Index(target, "("); p >= 0 {
		refTable = canonIdent(target[:p])
		if end := strings.LastIndex(target, ")"); end > p {
			refCols = parseColumnList(target[p+1 : end])
		}
	} else if consumed < len(toks) && strings.HasPrefix(toks[consumed], "(") {
		refCols = parseColumnList(innerParensFull(toks[consumed]))
		consumed++
	}
	fk := &ForeignKey{
		Name:       FKName(t.Name, localCols),
		Columns:    localCols,
		RefTable:   refTable,
		RefColumns: refCols,
	}
	for consumed+2 < len(toks) {
		if strings.ToUpper(toks[consumed]) != "ON" {
			break
		}
		action := strings.ToUpper(toks[consumed+1])
		val, step := collectOnAction(toks[consumed+2:])
		switch action {
		case "DELETE":
			fk.OnDelete = val
		case "UPDATE":
			fk.OnUpdate = val
		}
		consumed += 2 + step
	}
	return fk
}

func parseTableCheck(t *Table, name, item string) *CheckConstraint {
	expr := innerParens(afterKeyword(item, "CHECK"))
	if name == "" {
		name = strings.ToLower(t.Name) + "_check"
	}
	return &CheckConstraint{Name: name, Expression: strings.TrimSpace(expr)}
}

func parseNamedConstraint(t *Table, item string) {
	toks := tokenizeColumnDef(item)
	if len(toks) < 3 {
		return
	}
	name := canonIdent(toks[1])
	body := strings.Join(toks[2:], " ")
	upper := strings.ToUpper(body)
	switch {
	case strings.HasPrefix(upper, "PRIMARY KEY"):
		t.PrimaryKey = parseColumnList(afterKeyword(body, "PRIMARY KEY"))
	case strings.HasPrefix(upper, "UNIQUE"):
		cols := parseColumnList(afterKeyword(body, "UNIQUE"))
		t.Indexes = append(t.Indexes, &Index{Name: name, Columns: cols, Unique: true})
	case strings.HasPrefix(upper, "FOREIGN KEY"):
		fk := parseTableFK(t, body)
		if fk != nil {
			fk.Name = name
			t.ForeignKeys = append(t.ForeignKeys, fk)
		}
	case strings.HasPrefix(upper, "CHECK"):
		t.Checks = append(t.Checks, parseTableCheck(t, name, body))
	}
}

var reCreateIndex = regexp.MustCompile(`(?is)^\s*CREATE\s+(UNIQUE\s+)?INDEX\s+(?:IF\s+NOT\s+EXISTS\s+)?(\S+)\s+ON\s+(\S+?)(?:\s+USING\s+\S+)?\s*\((.*?)\)(?:\s+WHERE\s+(.*))?\s*$`)

func parseCreateIndex(s *Schema, stmt string) error {
	m := reCreateIndex.FindStringSubmatch(stmt)
	if m == nil {
		return nil // skip unparseable (permissive)
	}
	unique := strings.TrimSpace(m[1]) != ""
	name := canonIdent(m[2])
	tableName := canonIdent(m[3])
	cols := parseColumnList(m[4])
	where := strings.TrimSpace(m[5])
	t, ok := s.Tables[tableName]
	if !ok {
		// index references unknown table — create stub table so diff
		// sees it (real diff will flag it elsewhere).
		t = &Table{Name: tableName}
		s.Tables[tableName] = t
	}
	t.Indexes = append(t.Indexes, &Index{Name: name, Columns: cols, Unique: unique, Where: where})
	return nil
}

var reAlterAddFK = regexp.MustCompile(`(?is)^\s*ALTER\s+TABLE\s+(?:ONLY\s+)?(\S+)\s+ADD\s+(?:CONSTRAINT\s+(\S+)\s+)?FOREIGN\s+KEY\s*\(([^)]*)\)\s*REFERENCES\s+(\S+?)\s*\(([^)]*)\)(.*)$`)

func parseAlterTable(s *Schema, stmt string) error {
	m := reAlterAddFK.FindStringSubmatch(stmt)
	if m == nil {
		return nil
	}
	table := canonIdent(m[1])
	name := ""
	if m[2] != "" {
		name = canonIdent(m[2])
	}
	localCols := parseColumnList(m[3])
	refTable := canonIdent(m[4])
	refCols := parseColumnList(m[5])
	tail := strings.ToUpper(m[6])
	fk := &ForeignKey{Columns: localCols, RefTable: refTable, RefColumns: refCols}
	if name == "" {
		name = FKName(table, localCols)
	}
	fk.Name = name
	if strings.Contains(tail, "ON DELETE CASCADE") {
		fk.OnDelete = "CASCADE"
	} else if strings.Contains(tail, "ON DELETE SET NULL") {
		fk.OnDelete = "SET NULL"
	} else if strings.Contains(tail, "ON DELETE RESTRICT") {
		fk.OnDelete = "RESTRICT"
	} else if strings.Contains(tail, "ON DELETE NO ACTION") {
		fk.OnDelete = "NO ACTION"
	}
	if strings.Contains(tail, "ON UPDATE CASCADE") {
		fk.OnUpdate = "CASCADE"
	} else if strings.Contains(tail, "ON UPDATE SET NULL") {
		fk.OnUpdate = "SET NULL"
	} else if strings.Contains(tail, "ON UPDATE RESTRICT") {
		fk.OnUpdate = "RESTRICT"
	}
	t, ok := s.Tables[table]
	if !ok {
		t = &Table{Name: table}
		s.Tables[table] = t
	}
	t.ForeignKeys = append(t.ForeignKeys, fk)
	return nil
}

// splitTopLevel splits s on sep at top paren-depth zero.
func splitTopLevel(s string, sep byte) []string {
	var out []string
	var sb strings.Builder
	depth := 0
	inSQ, inDQ := false, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case inSQ:
			sb.WriteByte(c)
			if c == '\'' {
				if i+1 < len(s) && s[i+1] == '\'' {
					sb.WriteByte('\'')
					i++
				} else {
					inSQ = false
				}
			}
		case inDQ:
			sb.WriteByte(c)
			if c == '"' {
				inDQ = false
			}
		case c == '\'':
			inSQ = true
			sb.WriteByte(c)
		case c == '"':
			inDQ = true
			sb.WriteByte(c)
		case c == '(':
			depth++
			sb.WriteByte(c)
		case c == ')':
			depth--
			sb.WriteByte(c)
		case c == sep && depth == 0:
			out = append(out, sb.String())
			sb.Reset()
		default:
			sb.WriteByte(c)
		}
	}
	if sb.Len() > 0 {
		out = append(out, sb.String())
	}
	return out
}

func parseColumnList(s string) []string {
	// strip surrounding parens if present.
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		s = s[1 : len(s)-1]
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			continue
		}
		out = append(out, canonIdent(p))
	}
	return out
}

func innerParens(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		return strings.TrimSpace(s[1 : len(s)-1])
	}
	return s
}

// innerParensFull walks until matching paren (handles nested).
func innerParensFull(s string) string {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "(") {
		return s
	}
	depth := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			depth++
		} else if s[i] == ')' {
			depth--
			if depth == 0 {
				return s[1:i]
			}
		}
	}
	return s[1:]
}

// canonIdent returns the lowercase form of a PostgreSQL identifier,
// preserving the exact casing for "quoted" identifiers.
func canonIdent(s string) string {
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, ",")
	if strings.HasPrefix(s, `"`) && strings.HasSuffix(s, `"`) && len(s) >= 2 {
		return s[1 : len(s)-1]
	}
	return strings.ToLower(s)
}

func trimForErr(s string) string {
	s = strings.ReplaceAll(s, "\n", " ")
	if len(s) > 80 {
		return s[:80] + "..."
	}
	return s
}
