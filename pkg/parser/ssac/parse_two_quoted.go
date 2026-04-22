//ff:func feature=ssac-parse type=util control=sequence
//ff:what "first" "second" 형식의 두 개 인용문 파싱
package ssac

import "strings"

// parseTwoQuoted는 "first" "second"를 파싱한다.
func parseTwoQuoted(s string) (string, string, string) {
	s = strings.TrimSpace(s)
	first, rest := extractQuoted(s)
	second, remainder := extractQuoted(rest)
	return first, second, strings.TrimSpace(remainder)
}
