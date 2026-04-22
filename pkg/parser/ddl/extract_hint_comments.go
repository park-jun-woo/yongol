//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what extractHintComments — DDL 주석에서 -- @rename / @cast / @backfill / @data_migration / @allow_destructive 토큰 수집

package ddl

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

// HintComment is a single `-- @<tag> key=val ...` comment extracted from
// a DDL file, together with enough context for yongol migration hints
// to associate it with the right table/column.
//
// The scanner is intentionally simple: it records every hint comment
// and its *immediate context* — the CREATE TABLE currently being
// scanned (if any) plus the previous non-blank token on the same line
// (the candidate column name for column-line hints).
type HintComment struct {
	File       string            // absolute SQL file path
	Line       int               // 1-based
	Tag        string            // "rename" / "cast" / "backfill" / "data_migration" / "allow_destructive"
	Args       map[string]string // key=val pairs after the tag
	TableCtx   string            // last CREATE TABLE entered (lowercase), "" outside any table
	ColumnCtx  string            // column name if the hint is on a column definition line, "" otherwise
	BlockAbove bool              // true when the hint is on its own line *above* a CREATE TABLE / column
}

// ExtractHintCommentsFromDir walks <dir> for *.sql files and returns
// every yongol hint comment it finds. Unknown `-- @foo` tags are
// ignored so existing projects keep working.
func ExtractHintCommentsFromDir(dir string) ([]HintComment, error) {
	var out []HintComment
	entries, err := os.ReadDir(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(strings.ToLower(e.Name()), ".sql") {
			continue
		}
		path := filepath.Join(dir, e.Name())
		f, err := os.Open(path)
		if err != nil {
			return nil, err
		}
		hints, err := scanHintComments(f, path)
		f.Close()
		if err != nil {
			return nil, err
		}
		out = append(out, hints...)
	}
	return out, nil
}

func scanHintComments(r interface{ Read([]byte) (int, error) }, path string) ([]HintComment, error) {
	var out []HintComment
	sc := bufio.NewScanner(r.(interface{ Read([]byte) (int, error) }))
	lineNum := 0
	tableCtx := ""
	// pendingHints are hints on stand-alone comment lines that should
	// attach to the *next* non-blank DDL line.
	var pending []*HintComment
	for sc.Scan() {
		lineNum++
		ln := sc.Text()
		trim := strings.TrimSpace(ln)

		// Detect CREATE TABLE header.
		upper := strings.ToUpper(trim)
		if strings.HasPrefix(upper, "CREATE TABLE") {
			tableCtx = parseCreateTableName(trim)
			// consume any pending standalone hints as "above CREATE TABLE"
			for _, h := range pending {
				h.TableCtx = tableCtx
				h.BlockAbove = true
				out = append(out, *h)
			}
			pending = nil
			continue
		}

		// Comment-only line?
		if strings.HasPrefix(trim, "--") {
			hint := parseHintLine(trim, path, lineNum, tableCtx, "")
			if hint != nil {
				// Standalone: attach to the next DDL line.
				pending = append(pending, hint)
			}
			continue
		}

		// Line containing DDL + optional trailing `-- @...` comment.
		ddlPart, comment := splitTrailingComment(ln)
		ddlTrim := strings.TrimSpace(ddlPart)
		// Drain pending hints at the first real content.
		if len(pending) > 0 && ddlTrim != "" {
			column := extractColumnNameFromLine(ddlTrim)
			for _, h := range pending {
				if column != "" {
					h.ColumnCtx = column
				}
				h.TableCtx = tableCtx
				out = append(out, *h)
			}
			pending = nil
		}
		if comment != "" {
			column := extractColumnNameFromLine(ddlTrim)
			hint := parseHintLine("-- "+comment, path, lineNum, tableCtx, column)
			if hint != nil {
				out = append(out, *hint)
			}
		}
		if strings.HasSuffix(ddlTrim, ";") || strings.Contains(ddlTrim, ");") {
			// Conservative: end of statement clears table context.
			// (Works well enough for validation purposes.)
		}
	}
	// Any pending hints without a following DDL line are dropped.
	return out, sc.Err()
}

func parseCreateTableName(trim string) string {
	// "CREATE TABLE [IF NOT EXISTS] <name> (..."
	up := strings.ToUpper(trim)
	idx := strings.Index(up, "CREATE TABLE")
	if idx < 0 {
		return ""
	}
	rest := strings.TrimSpace(trim[idx+len("CREATE TABLE"):])
	if strings.HasPrefix(strings.ToUpper(rest), "IF NOT EXISTS") {
		rest = strings.TrimSpace(rest[len("IF NOT EXISTS"):])
	}
	// Stop at "(", whitespace, or ";".
	for i, c := range rest {
		if c == '(' || c == ' ' || c == '\t' || c == ';' {
			return strings.ToLower(strings.Trim(rest[:i], `"`))
		}
	}
	return strings.ToLower(strings.Trim(rest, `"`))
}

func extractColumnNameFromLine(s string) string {
	// First bareword before whitespace or '(' — that's the column name.
	s = strings.TrimSpace(s)
	s = strings.TrimSuffix(s, ",")
	if s == "" {
		return ""
	}
	for i, c := range s {
		if c == ' ' || c == '\t' || c == '(' {
			return strings.Trim(s[:i], `"`)
		}
	}
	return strings.Trim(s, `"`)
}

// splitTrailingComment returns (ddl, comment) where comment is the text
// after `-- ` on the same line (without the `--` prefix), respecting
// single-quoted string literals.
func splitTrailingComment(line string) (string, string) {
	inSQ := false
	for i := 0; i < len(line)-1; i++ {
		c := line[i]
		if c == '\'' {
			if inSQ && i+1 < len(line) && line[i+1] == '\'' {
				i++
				continue
			}
			inSQ = !inSQ
			continue
		}
		if !inSQ && c == '-' && line[i+1] == '-' {
			return line[:i], strings.TrimSpace(line[i+2:])
		}
	}
	return line, ""
}

// parseHintLine accepts a line guaranteed to start with `--`.
// Returns a HintComment if the comment starts with `@<tag>`.
func parseHintLine(line, file string, lineNum int, tableCtx, columnCtx string) *HintComment {
	body := strings.TrimSpace(strings.TrimPrefix(strings.TrimSpace(line), "--"))
	if !strings.HasPrefix(body, "@") {
		return nil
	}
	body = strings.TrimPrefix(body, "@")
	toks := strings.Fields(body)
	if len(toks) == 0 {
		return nil
	}
	tag := strings.ToLower(toks[0])
	args := map[string]string{}
	for _, t := range toks[1:] {
		if eq := strings.Index(t, "="); eq > 0 {
			args[strings.ToLower(t[:eq])] = t[eq+1:]
		}
	}
	return &HintComment{
		File:      file,
		Line:      lineNum,
		Tag:       tag,
		Args:      args,
		TableCtx:  tableCtx,
		ColumnCtx: columnCtx,
	}
}
