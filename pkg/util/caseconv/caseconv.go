//ff:func feature=util type=util control=iteration dimension=1 topic=string-convert
//ff:what SnakeToPascal — snake_case → PascalCase (plain: capitalize-first per part)

package caseconv

import (
	"strings"
)

// SnakeToPascal converts snake_case to PascalCase using the plain convention
// (capitalize-first per part). Example: "user_id" → "UserId", "per_page" → "PerPage".
func SnakeToPascal(s string) string {
	parts := strings.Split(s, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]))
		if len(p) > 1 {
			b.WriteString(p[1:])
		}
	}
	return b.String()
}
