//ff:func feature=gen-fastapi type=util control=sequence
//ff:what schemaPascalCase — camelCase/snake_case 문자열을 PascalCase 로 변환

package fastapi

import "strings"

// schemaPascalCase converts a camelCase or snake_case string to PascalCase.
func schemaPascalCase(s string) string {
	if s == "" {
		return ""
	}
	// If already PascalCase (starts with upper), return as-is.
	if s[0] >= 'A' && s[0] <= 'Z' {
		return s
	}
	// Convert first letter to uppercase for camelCase inputs.
	return strings.ToUpper(s[:1]) + s[1:]
}
