//ff:func feature=validate type=rule control=sequence topic=manifest-auth
//ff:what SEC-202 — auth.mode=bearer with csrf.enabled=false loses the runtime CSRF gate (BACKEND_AUTH_MODE=cookie|hybrid)

package manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// sec202RuntimeModeCsrf warns when a bearer-mode manifest explicitly
// disables CSRF (csrf.enabled: false).
//
// BUG-116 / Phase-B2 — the generated backend always carries the runtime
// authMode() switch (BACKEND_AUTH_MODE env), so cookie/hybrid transport
// is reachable at runtime even on a build whose manifest promises bearer.
// Phase-B1 closed the behavioral gap: the CSRF middleware is now emitted
// on every auth-enabled build and gated at runtime (no-op under bearer).
// The only configuration that still loses that runtime gate is an
// explicit csrf.enabled=false on a bearer manifest — the middleware is
// then not emitted at all, and an operator flipping
// BACKEND_AUTH_MODE=cookie|hybrid would authenticate via auto-attached
// cookies with no CSRF defense (the original BUG-116 hole, reopened by
// opt-out).
//
// Level is WARNING, not ERROR: the build-time default path (bearer mode,
// env unset) is safe, and the risk materializes only through an operator
// action. cookie/hybrid manifests with csrf disabled are SEC-201's ERROR
// domain; this rule covers exactly the bearer side SEC-201 cannot see.
// A jwt-typed manifest with mode unspecified resolves to "cookie" under
// ResolvedMode() and is therefore already rejected by SEC-201 — no
// overlap, no gap.
func sec202RuntimeModeCsrf(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.Manifest == nil || fs.Manifest.Backend.Auth == nil {
		return nil
	}
	a := fs.Manifest.Backend.Auth
	if a.ResolvedMode() != "bearer" {
		// cookie/hybrid (including the unspecified-mode default) with
		// csrf disabled is SEC-201's domain.
		return nil
	}
	if a.Csrf == nil || a.Csrf.Enabled {
		// Runtime-gated CSRF middleware is emitted (Phase-B1); the
		// build-time/runtime contract holds for every reachable mode.
		return nil
	}
	return []diagnostic.Diagnostic{{
		File:    "manifest.yaml",
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: "[SEC-202] auth.mode=\"bearer\" with csrf.enabled=false removes the runtime CSRF gate — running this build with BACKEND_AUTH_MODE=cookie|hybrid would use cookie auth with no CSRF defense",
		Advice:  "Drop csrf.enabled: false (or set it to true) so the runtime-gated CSRF middleware is emitted (no-op under bearer); keep it disabled only if BACKEND_AUTH_MODE is never switched to cookie/hybrid in operation",
	}}
}
