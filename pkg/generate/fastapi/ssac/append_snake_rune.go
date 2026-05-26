//ff:func feature=gen-fastapi type=util control=sequence
//ff:what appendSnakeRune — snake_case 변환용 개별 문자 추가 (대문자 앞 밑줄 삽입)

package ssac

import (
	"strings"
	"unicode"
)

// appendSnakeRune appends one rune to the snake_case builder, inserting
// an underscore before uppercase letters (except at position 0).
func appendSnakeRune(b *strings.Builder, i int, r rune) {
	if unicode.IsUpper(r) && i > 0 {
		b.WriteByte('_')
	}
	b.WriteRune(unicode.ToLower(r))
}
