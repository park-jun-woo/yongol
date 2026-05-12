//ff:func feature=gen-react type=util control=sequence
//ff:what kebab-case 문자열에 단순 영어 복수형 접미사를 추가한다

package react

import "strings"

// naivePluralize adds "s" to a kebab-case string for simple English
// pluralization. Handles common suffixes: -s, -sh, -ch, -x, -z → "es".
func naivePluralize(s string) string {
	if s == "" {
		return s
	}
	if strings.HasSuffix(s, "s") || strings.HasSuffix(s, "sh") ||
		strings.HasSuffix(s, "ch") || strings.HasSuffix(s, "x") ||
		strings.HasSuffix(s, "z") {
		return s + "es"
	}
	return s + "s"
}
