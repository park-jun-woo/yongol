//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what domainTitle — 도메인 이름을 PascalCase 접미사로 변환 (BearerAuthStrict<Title> 함수명 한정자)

package auth

import "strings"

// domainTitle returns the PascalCase form of domainIdent(name), used as the
// strict-middleware function-name suffix (BearerAuthStrictPublic). It mirrors
// gogin.domainTitle and boot.domainTitle; the auth package cannot import gogin
// (import cycle), so the rule is duplicated here. A single-site project passes
// "" and gets back "".
func domainTitle(name string) string {
	if name == "" {
		return ""
	}
	s := domainIdent(name)
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
