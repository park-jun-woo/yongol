//ff:func feature=migration type=util control=sequence
//ff:what afterKeyword — 문자열에서 키워드 이후 부분을 반환 (case-insensitive)
package migration

import "strings"

// afterKeyword returns the substring of s that comes after the first
// occurrence of kw (case-insensitive). Returns s unchanged when kw is
// not present.
func afterKeyword(s, kw string) string {
	idx := strings.Index(strings.ToUpper(s), strings.ToUpper(kw))
	if idx < 0 {
		return s
	}
	return strings.TrimSpace(s[idx+len(kw):])
}
