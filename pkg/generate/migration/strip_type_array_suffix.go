//ff:func feature=migration type=util control=iteration dimension=1
//ff:what stripTypeArraySuffix — 타입 문자열 끝의 `[]` 접미를 반복 제거 (array 여부 + trim 문자열)
package migration

import "strings"

// stripTypeArraySuffix returns (array, trimmed) where `array` is true
// when one or more `[]` suffixes were peeled off.
func stripTypeArraySuffix(s string) (bool, string) {
	array := false
	for strings.HasSuffix(s, "[]") {
		array = true
		s = strings.TrimSpace(strings.TrimSuffix(s, "[]"))
	}
	return array, s
}
