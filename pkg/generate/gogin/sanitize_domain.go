//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what sanitizeDomainName — 도메인 이름을 유효한 소문자 Go 패키지 식별자로 정규화
package gogin

// sanitizeDomainName converts a manifest domain key into a valid lowercase Go
// package identifier fragment. It lowercases the name, maps hyphens and any
// other non-identifier characters to underscores, and guards against a leading
// digit (Go identifiers may not begin with a digit) by prefixing an underscore.
// Example: "public" → "public", "my-admin" → "my_admin", "2nd" → "_2nd".
func sanitizeDomainName(name string) string {
	out := make([]rune, 0, len(name))
	for _, r := range name {
		switch {
		case r >= 'A' && r <= 'Z':
			out = append(out, r+('a'-'A'))
		case (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') || r == '_':
			out = append(out, r)
		default:
			out = append(out, '_')
		}
	}
	if len(out) == 0 {
		return "_"
	}
	if out[0] >= '0' && out[0] <= '9' {
		out = append([]rune{'_'}, out...)
	}
	return string(out)
}
