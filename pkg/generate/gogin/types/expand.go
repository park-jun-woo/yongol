//ff:func feature=gen-gogin type=util control=sequence
//ff:what Expand — GoTypeBinding 템플릿의 placeholder 치환 ({row}/{field}/{var})

package types

import "strings"

// Expand replaces the placeholders {row}, {field}, {var} in tmpl with the
// supplied values and returns the result. Placeholders that the template
// does not use are silently ignored.
//
// To emit a literal `{` write `{{`; to emit `}` write `}}`.
//
// Unknown placeholders are left intact so future template extensions stay
// non-breaking; today's grammar uses exactly the three above.
func Expand(tmpl, row, field, varName string) string {
	if tmpl == "" {
		return ""
	}
	// Two-pass replace: first protect literal escapes, then substitute,
	// then restore.
	const lbHolder = "\x00LBRACE\x00"
	const rbHolder = "\x00RBRACE\x00"
	s := strings.ReplaceAll(tmpl, "{{", lbHolder)
	s = strings.ReplaceAll(s, "}}", rbHolder)
	s = strings.ReplaceAll(s, "{row}", row)
	s = strings.ReplaceAll(s, "{field}", field)
	s = strings.ReplaceAll(s, "{var}", varName)
	s = strings.ReplaceAll(s, lbHolder, "{")
	s = strings.ReplaceAll(s, rbHolder, "}")
	return s
}
