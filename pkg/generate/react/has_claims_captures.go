//ff:func feature=gen-react type=util control=iteration dimension=2
//ff:what hasClaimsCaptures — 어느 페이지든 auth.claims.* 캡처가 선언되었는지 판정 (claims store 방출 게이트)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// hasClaimsCaptures reports whether any page's action declares an
// auth.claims.* data-capture — the gate for emitting the claims-carrying
// session store (plans/stml/sitemap Phase005). Projects without one keep
// the pre-Phase005 emission byte-identically ("claims only when used").
func hasClaimsCaptures(pages []stml.PageSpec) bool {
	for _, p := range pages {
		for _, a := range p.Actions {
			if actionHasClaimsCapture(a) {
				return true
			}
		}
	}
	return false
}
