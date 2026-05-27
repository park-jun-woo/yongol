//ff:func feature=gen-fastapi type=util control=iteration dimension=1
//ff:what snakeCase — PascalCase/camelCase → snake_case 변환 (연속 대문자 약어 올바르게 처리)

package ssac

import (
	"strings"
	"unicode"
)

// snakeCase converts a PascalCase or camelCase identifier to snake_case.
// Consecutive uppercase runs (e.g. "ID", "URL", "OrgID") are kept together:
//   - "ID" → "id", "OrgID" → "org_id", "ResolveRootID" → "resolve_root_id"
func snakeCase(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	var b strings.Builder
	for i, r := range runes {
		prevUpper := i > 0 && unicode.IsUpper(runes[i-1])
		nextLower := i+1 < len(runes) && unicode.IsLower(runes[i+1])
		appendSnakeRune(&b, i, r, prevUpper, nextLower)
	}
	return b.String()
}
