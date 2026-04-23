//ff:func feature=migration type=util control=sequence
//ff:what innerParens — 바깥쪽 "()" 를 제거해 내부 문자열만 반환
package migration

import "strings"

// innerParens returns the body of a parenthesised string. If s is not
// wrapped in parens, returns s unchanged.
func innerParens(s string) string {
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "(") && strings.HasSuffix(s, ")") {
		return strings.TrimSpace(s[1 : len(s)-1])
	}
	return s
}
