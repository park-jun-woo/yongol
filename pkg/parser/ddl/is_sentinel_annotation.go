//ff:func feature=manifest type=parser control=sequence
//ff:what isSentinelAnnotation — "-- @sentinel" 주석 라인 여부 확인
package ddl

import "strings"

// isSentinelAnnotation reports whether a trimmed line is a standalone
// `-- @sentinel` SQL comment.
func isSentinelAnnotation(trimmed string) bool {
	if !strings.HasPrefix(trimmed, "--") {
		return false
	}
	body := strings.TrimSpace(strings.TrimPrefix(trimmed, "--"))
	return body == "@sentinel"
}
