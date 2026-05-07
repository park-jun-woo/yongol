//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-sqlc
//ff:what pascalCaseSnake — snake_case 문자열을 언더스코어 기준 PascalCase로 조합

package ssac_sqlc

import "strings"

func pascalCaseSnake(s string) string {
	parts := strings.Split(s, "_")
	var b strings.Builder
	for _, p := range parts {
		if p == "" {
			continue
		}
		b.WriteString(strings.ToUpper(p[:1]) + p[1:])
	}
	return b.String()
}
