//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what isNumericLiteral -- 문자열이 숫자 리터럴(선택적 음수/소수점)인지 판정

package ir

// isNumericLiteral returns true when s is a numeric literal: optional leading
// minus, one or more digits, optional single decimal point with digits.
func isNumericLiteral(s string) bool {
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
	hasDot := false
	for i := start; i < len(s); i++ {
		if s[i] == '.' && !hasDot {
			hasDot = true
			continue
		}
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}
