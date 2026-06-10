//ff:func feature=gen-gogin type=util control=sequence topic=csrf
//ff:what hasCsrf — auth 선언 && csrf 미비활성 시 CSRF 블록 마운트 (BUG-116: bearer 빌드 포함, 런타임 게이트)

package boot

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// hasCsrf returns true when the generated server must mount the CSRF
// middleware. Active whenever auth is declared (and CSRF is not explicitly
// disabled), independent of the build-time resolved mode.
//
// Takes a resolved prepared.Auth so Mode is already defaulted; the raw
// vs resolved inconsistency that caused BUG-009 is structurally
// impossible here.
//
// BUG-116 / Phase-B1 — previously the block was gated on
// prepared.Auth.CsrfRequired (cookie/hybrid only), so a manifest=bearer
// build emitted no CSRF block. But the generated authMode() switch lets
// BACKEND_AUTH_MODE flip that same binary to cookie/hybrid at runtime,
// authenticating via cookies with no CSRF defense. The block is now always
// mounted; the emitted middleware no-ops at runtime in bearer mode via
// csrfRuntimeActive (see csrf_source.go), so bearer deployments are
// unchanged while the runtime-cookie gap is closed.
func hasCsrf(a prepared.Auth) bool {
	if !a.Present || a.Raw == nil {
		return false
	}
	if a.Raw.Csrf == nil {
		return true // auth present with unset csrf → enabled by default
	}
	return a.Raw.Csrf.Enabled
}
