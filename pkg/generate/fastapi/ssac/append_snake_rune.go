//ff:func feature=gen-fastapi type=util control=sequence
//ff:what appendSnakeRune — snake_case 변환용 개별 문자 추가 (연속 대문자 약어 올바르게 처리)

package ssac

import (
	"strings"
	"unicode"
)

// appendSnakeRune appends one rune to the snake_case builder, inserting
// an underscore before uppercase-to-lowercase transitions but not between
// consecutive uppercase letters. This ensures "OrgID" → "org_id" and
// "ResolveRootID" → "resolve_root_id" instead of "resolve_root_i_d".
//
// prevUpper indicates whether the rune at position i-1 was uppercase.
// nextLower indicates whether the rune at position i+1 is lowercase.
func appendSnakeRune(b *strings.Builder, i int, r rune, prevUpper, nextLower bool) {
	if unicode.IsUpper(r) && i > 0 {
		if !prevUpper || nextLower {
			b.WriteByte('_')
		}
	}
	b.WriteRune(unicode.ToLower(r))
}
