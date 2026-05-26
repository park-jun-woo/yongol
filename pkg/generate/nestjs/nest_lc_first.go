//ff:func feature=gen-nestjs type=util control=sequence
//ff:what nestLcFirst — 문자열 첫 글자 소문자 변환

package nestjs

import "unicode"

// nestLcFirst lowercases the first character of s.
func nestLcFirst(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToLower(r[0])
	return string(r)
}
