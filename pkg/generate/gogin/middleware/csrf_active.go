//ff:func feature=gen-gogin type=util control=sequence topic=csrf
//ff:what csrfActive — auth 선언 && csrf 미비활성 시 csrf.go 방출 (BUG-116: bearer 빌드 포함, 런타임 게이트)

package middleware

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// csrfActive mirrors the boot.hasCsrf gate so internal/middleware/csrf.go
// is emitted on the same trigger as the main.go registration. Takes a
// resolved prepared.Auth so the "raw vs resolved Mode" split that caused
// BUG-009 is structurally impossible: Mode here is always the defaulted
// value.
//
// BUG-116 / Phase-B1 — emission is gated on auth presence, not the
// build-time resolved mode. The emitted csrf.go is self-contained and
// no-ops at runtime in bearer mode (csrfRuntimeActive), so a manifest=bearer
// build still ships the CSRF defense for the case where BACKEND_AUTH_MODE
// flips it to cookie/hybrid at runtime.
func csrfActive(a prepared.Auth) bool {
	if !a.Present || a.Raw == nil {
		return false
	}
	if a.Raw.Csrf == nil {
		// Default on whenever auth is present — SEC-201 rejects the
		// explicit-false combination at validate time; reaching codegen
		// with nil Csrf means "accept defaults, enabled".
		return true
	}
	return a.Raw.Csrf.Enabled
}
