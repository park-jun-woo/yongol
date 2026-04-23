//ff:func feature=gen-gogin type=util control=sequence
//ff:what lowerFirst — 문자열의 첫 글자를 소문자로 (api field 이름 → 로컬 변수명)

package ssac

import "strings"

// lowerFirst lower-cases the leading rune. Used to derive a local
// variable name from an exported api field name ("PayloadTemplate" →
// "payloadTemplate"). The result is suffixed with "Map" at the call
// site so it clearly signals the decoded map.
func lowerFirst(s string) string {
	if s == "" {
		return s
	}
	return strings.ToLower(s[:1]) + s[1:]
}
