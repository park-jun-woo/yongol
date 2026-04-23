//ff:func feature=manifest type=parser control=iteration dimension=1
//ff:what extractColumnNameFromLine — DDL 라인에서 첫 번째 bareword 를 column 이름으로 반환

package ddl

import "strings"

// extractColumnNameFromLine returns the first bareword of a DDL line —
// the candidate column name. Handles quoted identifiers and trailing commas.
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
