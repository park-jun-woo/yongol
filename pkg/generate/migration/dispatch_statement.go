//ff:func feature=migration type=parser control=selection
//ff:what dispatchStatement — 한 SQL statement 를 CREATE TABLE / CREATE INDEX / ALTER TABLE 파서로 분기
package migration

import "strings"

// dispatchStatement routes a single trimmed SQL statement to the right
// parser. Unknown statements (INSERT, COMMENT, SET, …) are ignored.
func dispatchStatement(s *Schema, stmt string) error {
	if stmt == "" {
		return nil
	}
	upper := strings.ToUpper(stmt)
	switch {
	case strings.HasPrefix(upper, "CREATE TABLE"):
		return parseCreateTable(s, stmt)
	case strings.HasPrefix(upper, "CREATE UNIQUE INDEX"),
		strings.HasPrefix(upper, "CREATE INDEX"):
		return parseCreateIndex(s, stmt)
	case strings.HasPrefix(upper, "ALTER TABLE"):
		return parseAlterTable(s, stmt)
	}
	return nil
}
