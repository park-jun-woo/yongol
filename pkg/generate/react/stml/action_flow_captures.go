//ff:func feature=stml-gen type=util control=sequence
//ff:what ActionBlock의 유효 data-capture 바인딩 반환 — bearer 는 전부, cookie 는 auth.claims.* 만
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// actionFlowCaptures returns the action's effective data-capture bindings.
// In bearer mode every binding commits to the session store. In cookie
// mode (or without backend.auth) the token sinks are ignored — httpOnly
// cookies carry the session and TM-24 diagnoses the stale declarations —
// but auth.claims.* bindings stay effective (plans/stml/sitemap Phase005):
// a claim is captured from the login response *body*, not the cookie, so
// cookie mode commits it to the claims-only store the same way.
func actionFlowCaptures(a stmlparser.ActionBlock, bearerAuth bool) []stmlparser.CaptureBind {
	if bearerAuth {
		return a.Captures
	}
	return claimsCaptures(a.Captures)
}
