//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what isIntegerLiteralStr — 문자열이 정수 리터럴인지 판별 (선택적 음수 부호 + 숫자)

package ssac

import "unicode"

// isIntegerLiteralStr returns true when s is a plain integer literal
// (optional leading minus followed by digits only).
func isIntegerLiteralStr(s string) bool {
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
	for i := start; i < len(s); i++ {
		if !unicode.IsDigit(rune(s[i])) {
			return false
		}
	}
	return true
}
