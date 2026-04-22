//ff:func feature=validate type=util control=sequence topic=ddl-structural
//ff:what isSkippableDDLLine — 컬럼 정의가 아닌 스킵 대상 DDL 라인 여부 검사
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
