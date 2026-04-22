//ff:func feature=validate type=util control=sequence topic=ssac-structural
//ff:what s62prefix — `.` 이전 부분만 반환 (S-62 변수 참조 비교용)

package ssac

import (
	"strings"
)

// s62prefix returns the portion of s before the first '.', or s itself.
func s62prefix(s string) string {
	if dot := strings.IndexByte(s, '.'); dot > 0 {
		return s[:dot]
	}
	return s
}
