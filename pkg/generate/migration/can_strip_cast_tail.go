//ff:func feature=migration type=util control=sequence
//ff:what canStripCastTail — ::뒤 문자열이 안전하게 제거 가능한 단순 식별자인지 판정
package migration

import "strings"

// canStripCastTail reports whether a trailing ::<tail> can be safely
// stripped from a DEFAULT expression.
func canStripCastTail(tail string) bool {
	if tail == "" {
		return false
	}
	if strings.Count(tail, "(") != strings.Count(tail, ")") {
		return false
	}
	return looksLikeCastTarget(tail)
}
