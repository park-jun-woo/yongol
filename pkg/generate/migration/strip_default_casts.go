//ff:func feature=migration type=parser control=iteration dimension=1
//ff:what stripDefaultCasts — DEFAULT 표현식 끝의 ::type 캐스트를 반복 제거
package migration

import "strings"

// stripDefaultCasts removes trailing `::type` casts from a DEFAULT
// expression, but only when the cast target looks like a bare
// identifier (possibly quoted). Stops on unbalanced parens or
// non-identifier targets.
func stripDefaultCasts(s string) string {
	for {
		idx := strings.LastIndex(s, "::")
		if idx < 0 {
			break
		}
		tail := strings.TrimSpace(s[idx+2:])
		if !canStripCastTail(tail) {
			break
		}
		s = strings.TrimSpace(s[:idx])
	}
	return s
}
