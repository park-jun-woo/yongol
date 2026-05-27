//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what snakeCase — PascalCase/camelCase → snake_case 변환 (연속 대문자 약어 올바르게 처리)

package funcstub

import (
	"strings"
	"unicode"
)

// snakeCase converts PascalCase/camelCase to snake_case for Python function
// names. Consecutive uppercase runs (e.g. "ID", "URL") are kept together.
func snakeCase(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	var b strings.Builder
	for i, r := range runes {
		prevUpper := i > 0 && unicode.IsUpper(runes[i-1])
		nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
		needsUnderscore := unicode.IsUpper(r) && i > 0 && (!prevUpper || nextLower)
		if needsUnderscore {
			b.WriteByte('_')
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
