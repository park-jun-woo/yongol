//ff:func feature=stml-gen type=util control=sequence topic=string-convert
//ff:what 문자열의 첫 글자를 대문자로 변환한다
package stml

import "unicode"

func toUpperFirst(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}
