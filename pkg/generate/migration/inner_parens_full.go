//ff:func feature=migration type=util control=sequence
//ff:what innerParensFull — 첫 "(" 의 짝 맞는 ")" 까지 내부 반환 (중첩 괄호 지원)
package migration

import "strings"

// innerParensFull walks until the matching close paren, supporting
// nested groups. Returns everything between the outer parens.
func innerParensFull(s string) string {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, "(") {
		return s
	}
	if end := matchingCloseParen(s); end > 0 {
		return s[1:end]
	}
	return s[1:]
}
