//ff:func feature=gen-nestjs type=util control=iteration dimension=1
//ff:what toSnake — Go PascalCase 식별자 → snake_case 변환 (ID → id 특수 처리)

package ssac

import "strings"

// toSnake converts a Go-style identifier to snake_case for TypeScript/Prisma
// field access. Special-cases "ID" → "id".
func toSnake(s string) string {
	if s == "ID" || s == "Id" {
		return "id"
	}
	var result strings.Builder
	for i, r := range s {
		isUpper := r >= 'A' && r <= 'Z'
		needsUnderscore := isUpper && i > 0 && s[i-1] >= 'a' && s[i-1] <= 'z'
		if needsUnderscore {
			result.WriteByte('_')
		}
		if isUpper {
			result.WriteByte(byte(r - 'A' + 'a'))
		} else {
			result.WriteByte(byte(r))
		}
	}
	return result.String()
}
