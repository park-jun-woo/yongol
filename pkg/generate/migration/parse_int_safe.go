//ff:func feature=migration type=util control=iteration dimension=1
//ff:what parseIntSafe — 숫자만 있는 문자열을 int 로 파싱, 이외 문자면 0 반환
package migration

// parseIntSafe returns the integer value of s, or 0 when s contains any
// non-digit rune.
func parseIntSafe(s string) int {
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0
		}
		n = n*10 + int(r-'0')
	}
	return n
}
