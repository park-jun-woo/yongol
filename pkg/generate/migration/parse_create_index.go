//ff:func feature=migration type=parser control=sequence
//ff:what parseCreateIndex — CREATE [UNIQUE] INDEX 문장 파싱
package migration

import (
	"regexp"
	"strings"
)

var reCreateIndex = regexp.MustCompile(`(?is)^\s*CREATE\s+(UNIQUE\s+)?INDEX\s+(?:IF\s+NOT\s+EXISTS\s+)?(\S+)\s+ON\s+(\S+?)(?:\s+USING\s+(\S+))?\s*\((.*?)\)(?:\s+WHERE\s+(.*))?\s*$`)

// parseCreateIndex handles CREATE INDEX / CREATE UNIQUE INDEX. Unparseable
// statements are silently skipped (permissive).
func parseCreateIndex(s *Schema, stmt string) error {
	m := reCreateIndex.FindStringSubmatch(stmt)
	if m == nil {
		return nil
	}
	unique := strings.TrimSpace(m[1]) != ""
	name := canonIdent(m[2])
	tableName := canonIdent(m[3])
	method := strings.ToLower(strings.TrimSpace(m[4]))
	cols := parseColumnList(m[5])
	where := strings.TrimSpace(m[6])
	t := ensureTable(s, tableName)
	t.Indexes = append(t.Indexes, &Index{Name: name, Columns: cols, Unique: unique, Method: method, Where: where})
	return nil
}
