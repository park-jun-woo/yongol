//ff:func feature=rule type=util control=iteration dimension=1 topic=catalog
//ff:what sectionAnchor — H2 heading 을 GFM slug (소문자 + hyphen) 로 변환

package catalog

import "strings"

// sectionAnchor converts an H2 heading into a GFM-style slug:
// lowercase, dots/spaces → hyphen, punctuation stripped.
// Used for helpUri construction.
func sectionAnchor(title string) string {
	var b strings.Builder
	b.Grow(len(title))
	prevDash := false
	for _, r := range strings.ToLower(title) {
		switch {
		case r >= 'a' && r <= 'z', r >= '0' && r <= '9':
			b.WriteRune(r)
			prevDash = false
		case r == ' ', r == '-', r == '.', r == '/':
			if !prevDash && b.Len() > 0 {
				b.WriteByte('-')
				prevDash = true
			}
		}
	}
	out := b.String()
	return strings.TrimRight(out, "-")
}
