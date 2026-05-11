//ff:func feature=stml-parse type=util control=iteration dimension=1 topic=string-convert
//ff:what kebab-case 문자열을 camelCase로 변환
package stml

import "strings"

// kebabToCamel converts kebab-case to camelCase.
func kebabToCamel(s string) string {
	if !strings.Contains(s, "-") {
		return s
	}
	parts := strings.Split(s, "-")
	for i := 1; i < len(parts); i++ {
		if len(parts[i]) > 0 {
			parts[i] = strings.ToUpper(parts[i][:1]) + parts[i][1:]
		}
	}
	return strings.Join(parts, "")
}
