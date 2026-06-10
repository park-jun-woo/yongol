//ff:func feature=gen-ir type=util control=sequence
//ff:what csrfIsActive -- prepared.Auth 기반 CSRF 미들웨어 활성화 여부

package ir

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// csrfIsActive returns true when the CSRF middleware should be emitted.
//
// BUG-116 / Phase-B1 — the gate is auth *presence*, not the build-time
// resolved mode (ps.Auth.CsrfRequired). Whenever auth is declared the
// generator emits the runtime authMode() switch (BACKEND_AUTH_MODE), so
// cookie/hybrid transport is reachable at runtime even on a manifest=bearer
// build. Emitting CSRF unconditionally (and gating it at runtime inside the
// middleware via csrfRuntimeActive) closes the gap where a bearer build run
// as BACKEND_AUTH_MODE=cookie authenticated via cookies with no CSRF check.
// Explicit csrf.enabled=false still opts out.
func csrfIsActive(ps *prepared.State) bool {
	if !ps.Auth.Present || ps.Auth.Raw == nil {
		return false
	}
	if ps.Auth.Raw.Csrf == nil {
		return true // auth present with unset csrf -> enabled by default
	}
	return ps.Auth.Raw.Csrf.Enabled
}
