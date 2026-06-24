//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what domainIdent — 도메인 이름을 유효한 소문자 Go 식별자 조각으로 정규화 (openapi_<x>.yaml·embed 변수용)

package middleware

// domainIdent normalizes a manifest domain key into a valid lowercase Go
// identifier fragment, mirroring gogin.sanitizeDomainName (Phase006) and
// auth.domainIdent (Phase008): lowercase, non-identifier runes → '_',
// leading-digit guard. The middleware package cannot import gogin (gogin
// imports middleware — import cycle), so the rule is duplicated here and must
// stay in lockstep so the openapi_<ident>.yaml embed filename matches the
// api_<ident> package generateAPIPerDomain creates.
func domainIdent(name string) string {
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
