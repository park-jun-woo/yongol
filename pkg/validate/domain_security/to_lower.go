//ff:func feature=validate type=util control=iteration dimension=1 topic=domain-security
//ff:what toLower — ASCII 소문자 변환 (strings import 회피용)
package domain_security

// toLower is a simple ASCII lowercase helper to avoid importing strings.
func toLower(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + 32
		}
	}
	return string(b)
}
