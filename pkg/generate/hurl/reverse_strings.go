//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what reverseStrings — 문자열 슬라이스 역순 반환
package hurl

// reverseStrings returns a reversed copy of the input slice.
func reverseStrings(s []string) []string {
	result := make([]string, len(s))
	for i, v := range s {
		result[len(s)-1-i] = v
	}
	return result
}
