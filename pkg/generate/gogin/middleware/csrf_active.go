//ff:func feature=gen-gogin type=util control=sequence topic=csrf
//ff:what csrfActive — prepared.Auth.Mode=cookie|hybrid 이고 csrf.enabled=true 인지 판정

package middleware

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// csrfActive mirrors the boot.blockCsrf gate so the middleware file is
// emitted on the same trigger. Takes a resolved prepared.Auth so the
// "raw vs resolved Mode" split that caused BUG-009 is structurally
// impossible: Mode here is always the defaulted value.
//
// Phase001 debug01 (BUG-013) — the "must emit CSRF?" decision is
// now carried by prepared.Auth.CsrfRequired. csrfActive short-circuits
// on that flag so JWT-only projects never produce csrf.go.
func csrfActive(a prepared.Auth) bool {
	if !a.CsrfRequired {
		return false
	}
	if !a.Present || a.Raw == nil {
		return false
	}
	if a.Raw.Csrf == nil {
		// Default on for cookie/hybrid modes — SEC-201 rejects the
		// explicit-false combination at validate time; reaching codegen
		// with nil Csrf means "accept defaults, enabled".
		return true
	}
	return a.Raw.Csrf.Enabled
}
