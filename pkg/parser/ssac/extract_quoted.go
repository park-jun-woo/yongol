//ff:func feature=ssac-parse type=util control=sequence
//ff:what 문자열 앞의 "quoted" 값을 추출하고 나머지 반환
package ssac

import "strings"

// extractQuoted는 문자열 앞의 "quoted" 값을 추출하고 나머지를 반환한다.
func extractQuoted(s string) (string, string) {
	s = strings.TrimSpace(s)
	if !strings.HasPrefix(s, `"`) {
		return "", s
	}
	endIdx := strings.IndexByte(s[1:], '"')
	if endIdx < 0 {
		return "", s
	}
	return s[1 : endIdx+1], strings.TrimSpace(s[endIdx+2:])
}
