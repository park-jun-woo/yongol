//ff:func feature=ssac-parse type=util control=sequence
//ff:what "target \"message\"" 문자열을 target과 message로 분리
package ssac

import "strings"

// splitTargetMessage는 "target "message""를 분리한다.
func splitTargetMessage(s string) (string, string, string) {
	quoteIdx := strings.IndexByte(s, '"')
	if quoteIdx < 0 {
		return strings.TrimSpace(s), "", ""
	}
	target := strings.TrimSpace(s[:quoteIdx])
	msg, remainder := extractQuoted(s[quoteIdx:])
	return target, msg, strings.TrimSpace(remainder)
}
