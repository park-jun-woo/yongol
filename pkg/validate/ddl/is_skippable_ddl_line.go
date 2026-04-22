//ff:func feature=validate type=util control=sequence topic=ddl-structural
//ff:what isSkippableDDLLine — report whether a DDL line is not a column definition and should be skipped
package ddl

import "strings"

// isSkippableDDLLine returns true for lines that aren't column definitions.
func isSkippableDDLLine(trimmed string) bool {
	if trimmed == "" || strings.HasPrefix(trimmed, "--") || strings.HasPrefix(trimmed, ")") {
		return true
	}
	upper := strings.ToUpper(trimmed)
	if strings.HasPrefix(upper, "CREATE") || strings.HasPrefix(upper, "INSERT") ||
		strings.HasPrefix(upper, "ON ") || strings.HasPrefix(upper, "VALUES") {
		return true
	}
	if strings.HasPrefix(upper, "PRIMARY KEY") || strings.HasPrefix(upper, "UNIQUE") ||
		strings.HasPrefix(upper, "CHECK") || strings.HasPrefix(upper, "FOREIGN KEY") ||
		strings.HasPrefix(upper, "CONSTRAINT") {
		return true
	}
	return false
}
