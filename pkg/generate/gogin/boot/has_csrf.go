//ff:func feature=gen-gogin type=util control=sequence topic=csrf
//ff:what hasCsrf — prepared.Auth.Mode=cookie|hybrid && csrf.enabled 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// hasCsrf returns true when the generated server must mount the CSRF
// middleware. Active only for cookie/hybrid auth modes with CSRF
// enabled. The default bearer mode is inherently CSRF-immune (attacker
// cannot inject the Authorization header cross-origin) so the middleware
// stays dormant — no overhead on the common path.
//
// Takes a resolved prepared.Auth so Mode is already defaulted; the raw
// vs resolved inconsistency that caused BUG-009 is structurally
// impossible here.
func hasCsrf(a prepared.Auth) bool {
	if !a.Present || a.Raw == nil {
		return false
	}
	if a.Mode != "cookie" && a.Mode != "hybrid" {
		return false
	}
	if a.Raw.Csrf == nil {
		return true // cookie/hybrid with unset csrf → enabled by default
	}
	return a.Raw.Csrf.Enabled
}
