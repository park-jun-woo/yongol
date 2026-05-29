//ff:func feature=manifest type=parser control=sequence
//ff:what isFuncManagedAnnotation — "-- @func-managed" 주석 라인 여부 확인
package ddl

import "strings"

// isFuncManagedAnnotation reports whether a trimmed line is a standalone
// `-- @func-managed` SQL comment. Such tables are actively managed by a
// `@call`'d function/RPC (not orphaned), so XSD-55 exempts them while other
// rules still apply.
func isFuncManagedAnnotation(trimmed string) bool {
	if !strings.HasPrefix(trimmed, "--") {
		return false
	}
	body := strings.TrimSpace(strings.TrimPrefix(trimmed, "--"))
	return body == "@func-managed"
}
