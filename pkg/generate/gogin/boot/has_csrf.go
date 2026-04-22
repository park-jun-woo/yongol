//ff:func feature=gen-gogin type=util control=sequence topic=csrf
//ff:what hasCsrf — manifest.backend.auth.mode != "bearer" && csrf.enabled 여부

package boot

import "github.com/park-jun-woo/yongol/pkg/yongol"

// hasCsrf returns true when the generated server must mount the CSRF
// middleware. Active only for cookie/hybrid auth modes with CSRF
// enabled. The default bearer mode is inherently CSRF-immune (attacker
// cannot inject the Authorization header cross-origin) so the middleware
// stays dormant — no overhead on the common path.
//
// Phase005 ships the wiring; Phase020 flips the default projects to
// cookie mode once the session-issuance pipeline is in place.
func hasCsrf(fs *yongol.Fullstack) bool {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return false
	}
	a := fs.Manifest.Backend.Auth
	// Phase020 — ResolvedMode() applies the "cookie" default so a
	// manifest with no backend.auth.mode set still lights up CSRF.
	mode := a.ResolvedMode()
	if mode != "cookie" && mode != "hybrid" {
		return false
	}
	if a.Csrf == nil {
		return true // cookie/hybrid with unset csrf → enabled by default
	}
	return a.Csrf.Enabled
}
