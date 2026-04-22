//ff:func feature=validate type=rule control=sequence topic=manifest-auth
//ff:what SEC-201 — auth.mode=cookie|hybrid combined with csrf.enabled=false is forbidden

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// sec201CookieWithoutCsrf rejects manifests that configure cookie-based
// session authentication (Mode == "cookie" or "hybrid") while explicitly
// disabling CSRF defense. Cookie-bound requests are vulnerable to CSRF
// by design — the browser auto-attaches the session cookie — so allowing
// the combination would silently ship an exploitable server.
//
// Note: Phase005 keeps the default bearer mode untouched, so this rule
// fires only for manifests that explicitly opt into cookie/hybrid modes
// ahead of Phase020's full session infrastructure.
func sec201CookieWithoutCsrf(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	a := fs.Manifest.Backend.Auth
	// Phase020 — ResolvedMode() applies the "cookie" default so a
	// manifest with no backend.auth.mode set is correctly gated.
	mode := a.ResolvedMode()
	if mode != "cookie" && mode != "hybrid" {
		return nil
	}
	if a.Csrf == nil {
		// nil == accept defaults (enabled). Only an explicit
		// csrf.enabled=false is flagged.
		return nil
	}
	if a.Csrf.Enabled {
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: "[SEC-201] auth.mode=\"" + mode + "\" requires csrf.enabled to be true",
		Advice:  "Set backend.auth.csrf.enabled to true, or change auth.mode to \"bearer\"",
	}}
}
