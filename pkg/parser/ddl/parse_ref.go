//ff:func feature=manifest type=util control=sequence
//ff:what parseRef — "users(id)" → (table, col) 파싱
package ddl

import "strings"

func parseRef(s string) (string, string) {
	s = strings.TrimSpace(s)
	parenIdx := strings.Index(s, "(")
	if parenIdx < 0 {
		return s, ""
	}
	table := s[:parenIdx]
	col := strings.TrimSuffix(s[parenIdx+1:], ")")
	col = strings.TrimSuffix(col, ",")
	return table, col
}
