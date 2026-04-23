//ff:func feature=rule type=util control=selection topic=catalog
//ff:what writeSectionAnchorRune — sectionAnchor 루프 본문. 1 rune 을 buffer 에 쓰고 새 prevDash 상태 반환

package catalog

import "strings"

func writeSectionAnchorRune(b *strings.Builder, r rune, prevDash bool) bool {
	switch {
	case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
		b.WriteRune(r)
		return false
	case r == ' ', r == '-', r == '.', r == '/':
		if prevDash || b.Len() == 0 {
			return prevDash
		}
		b.WriteByte('-')
		return true
	}
	return prevDash
}
