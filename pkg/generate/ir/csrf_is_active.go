//ff:func feature=gen-ir type=util control=sequence
//ff:what csrfIsActive -- prepared.Auth 기반 CSRF 미들웨어 활성화 여부

package ir

import "github.com/park-jun-woo/yongol/pkg/generate/prepared"

// csrfIsActive returns true when CSRF middleware should be active.
func csrfIsActive(ps *prepared.State) bool {
	if !ps.Auth.CsrfRequired {
		return false
	}
	if !ps.Auth.Present || ps.Auth.Raw == nil {
		return false
	}
	if ps.Auth.Raw.Csrf == nil {
		return true // cookie/hybrid with unset csrf -> enabled by default
	}
	return ps.Auth.Raw.Csrf.Enabled
}
