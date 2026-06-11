//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what hasClaimsCapture — 어느 페이지든 auth.claims.<name> 캡처가 선언되었는지 판정 (TM-47 배선 검사)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// hasClaimsCapture reports whether any page's action declares a
// data-capture into "auth.claims.<name>" — the wiring TM-47 requires for
// the sitemap data-roles menu filter. Sink matching goes through
// stml.ClaimsSinkName so the judgment cannot drift from the parser's.
func hasClaimsCapture(pages []stml.PageSpec, name string) bool {
	for _, c := range collectPageCaptures(pages) {
		if got, ok := stml.ClaimsSinkName(c.Bind.Sink); ok && got == name {
			return true
		}
	}
	return false
}
