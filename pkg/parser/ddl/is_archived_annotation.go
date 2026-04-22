//ff:func feature=manifest type=parser control=sequence
//ff:what isArchivedAnnotation — "-- @archived" 주석 라인 여부 확인
package ddl

import "strings"

// isArchivedAnnotation reports whether a trimmed line is a standalone
// `-- @archived` SQL comment.
func isArchivedAnnotation(trimmed string) bool {
	if !strings.HasPrefix(trimmed, "--") {
		return false
	}
	body := strings.TrimSpace(strings.TrimPrefix(trimmed, "--"))
	return body == "@archived"
}
