//ff:func feature=migration type=util control=sequence
//ff:what sortStringSlice — 문자열 슬라이스 제자리 정렬 (sort.Strings 얇은 래퍼)
package migration

import "sort"

// sortStringSlice sorts the given string slice in place.
func sortStringSlice(s []string) { sort.Strings(s) }
