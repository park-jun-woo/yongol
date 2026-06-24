//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what domainTitle — 도메인 이름을 PascalCase 식별자 조각으로 변환 (단일 사이트는 빈 문자열)
package gogin

import "strings"

// domainTitle returns the PascalCase form of sanitizeDomainName(name), used as
// a converter function-name prefix so two domains that share a schema name emit
// distinct convert<DomainTitle><Name> functions in the shared internal/service
// package (e.g. "public" → "Public", "my-admin" → "MyAdmin"). A single-site
// project passes the empty name and gets back "" — the degenerate case that
// keeps convert<Name> output unchanged.
func domainTitle(name string) string {
	if name == "" {
		return ""
	}
	s := sanitizeDomainName(name)
	var b strings.Builder
	upper := true
	for _, r := range s {
		if r == '_' {
			upper = true
			continue
		}
		if upper && r >= 'a' && r <= 'z' {
			r -= 'a' - 'A'
		}
		upper = false
		b.WriteRune(r)
	}
	return b.String()
}
