//ff:func feature=gen-gogin type=util control=sequence
//ff:what pascalCase — snake_case 또는 lowerCamel 을 PascalCase로 변환

package ssac

import "strings"

// pascalCase: "id" → "Id", "org_id" → "OrgId", "ID" → "ID".
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
