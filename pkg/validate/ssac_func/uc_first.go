//ff:func feature=validate type=util control=sequence topic=func-check
//ff:what ucFirst — 첫 글자를 대문자로 변환 (ASCII only, Go 식별자용)

package ssac_func

// ucFirst upper-cases the first byte of s. ASCII-only (sufficient for Go
// identifier names).
func ucFirst(s string) string {
	if s == "" {
		return s
	}
	if s[0] >= 'a' && s[0] <= 'z' {
		return string(s[0]-32) + s[1:]
	}
	return s
}
