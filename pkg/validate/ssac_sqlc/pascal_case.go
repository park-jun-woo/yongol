//ff:func feature=validate type=util control=sequence topic=ssac-sqlc
//ff:what pascalCase — snake_case 또는 lowerCamel 을 PascalCase로 변환 (codegen 동일 로직)

package ssac_sqlc

import "strings"

func pascalCase(s string) string {
	if s == "" {
		return s
	}
	if strings.Contains(s, "_") {
		return pascalCaseSnake(s)
	}
	if s[0] >= 'A' && s[0] <= 'Z' {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
