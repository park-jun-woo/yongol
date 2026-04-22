//ff:func feature=validate type=util control=sequence topic=ddl-structural
//ff:what hasSentinelInsert — report whether an INSERT INTO <table> VALUES (0, ...) sentinel exists
package ddl

import "regexp"

// hasSentinelInsert returns true if the SQL content contains
// `INSERT INTO <table> ... VALUES (0, ...)` for the given table.
func hasSentinelInsert(content, tableName string) bool {
	insertRe := regexp.MustCompile(`(?i)INSERT\s+INTO\s+` + regexp.QuoteMeta(tableName) + `\b`)
	loc := insertRe.FindStringIndex(content)
	if loc == nil {
		return false
	}
	valuesRe := regexp.MustCompile(`(?i)VALUES\s*\(\s*0\s*,`)
	return valuesRe.MatchString(content[loc[0]:])
}
