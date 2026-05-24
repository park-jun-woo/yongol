//ff:func feature=validate type=util control=iteration dimension=1 topic=ssac-structural
//ff:what xss60IsIntLiteral — 문자열이 정수 리터럴인지 판별 (선택적 음수 부호 + 숫자)

package ssac

import "unicode"

// xss60IsIntLiteral returns true if s consists entirely of digits, optionally
// preceded by a minus sign.
func xss60IsIntLiteral(s string) bool {
	if len(s) == 0 {
		return false
	}
	start := 0
	if s[0] == '-' {
		start = 1
	}
	if start >= len(s) {
		return false
	}
	for _, r := range s[start:] {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}
