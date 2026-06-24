//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what domainTitle — 도메인 이름을 PascalCase 접미사로 변환 (RequestValidator<Title> 함수명 한정자)

package middleware

import "strings"

// domainTitle returns the PascalCase form of domainIdent(name), used as the
// per-domain validator function-name suffix (RequestValidatorPublic) and embed
// variable suffix (openapiSpecPublic). It mirrors auth.domainTitle and
// boot.domainTitle; the middleware package cannot import gogin (import cycle),
// so the rule is duplicated here. A single-site project never calls this (the
// single-site validator keeps its unsuffixed RequestValidator name).
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
