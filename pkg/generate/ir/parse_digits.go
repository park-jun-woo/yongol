//ff:func feature=gen-ir type=util control=iteration dimension=1
//ff:what parseDigits -- 문자열에서 숫자만 추출하여 int64 반환 (비숫자 시 -1)

package ir

// parseDigits converts a string of ASCII digits to int64. Returns -1 if
// the string contains any non-digit character or is empty.
func parseDigits(s string) int64 {
	if s == "" {
		return -1
	}
	var n int64
	for _, c := range s {
		if c < '0' || c > '9' {
			return -1
		}
		n = n*10 + int64(c-'0')
	}
	return n
}
