//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what claimsCaptures — 캡처 바인딩에서 auth.claims.* sink 만 골라 반환

package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// claimsCaptures filters the bindings down to the auth.claims.* sinks —
// the only captures that stay effective in cookie mode (plans/stml/sitemap
// Phase005). Sink matching goes through stmlparser.ClaimsSinkName so the
// judgment cannot drift from the parser's.
func claimsCaptures(binds []stmlparser.CaptureBind) []stmlparser.CaptureBind {
	var claims []stmlparser.CaptureBind
	for _, c := range binds {
		if _, ok := stmlparser.ClaimsSinkName(c.Sink); ok {
			claims = append(claims, c)
		}
	}
	return claims
}
