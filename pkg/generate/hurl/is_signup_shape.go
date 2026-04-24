//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what isSignupShape — SSaC ServiceFunc 가 auth.HashPassword 호출을 포함하는지 판정

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// isSignupShape returns true when the service func contains an
// `@call auth.HashPassword(...)` sequence. The built-in package
// `github.com/park-jun-woo/ssac/pkg/auth` exposes HashPassword with
// that exact spelling, so presence of the call is a sufficient
// signal for a signup endpoint (even when the project wraps it via
// a local func/*.ssac override — the target symbol stays "auth.HashPassword").
//
// Projects commonly pair this with `@post <Model>.Create({..., PasswordHash: ...})`
// — that pairing is NOT required for detection (heuristic kept simple to
// avoid false negatives on combined or atypical signup flows). Callers
// may log a WARNING when the `@post User.Create` companion is missing.
func isSignupShape(fn *ssac.ServiceFunc) bool {
	if fn == nil {
		return false
	}
	for _, seq := range fn.Sequences {
		if seq.Type != ssac.SeqCall {
			continue
		}
		if matchesHashPassword(seq.Model) {
			return true
		}
	}
	return false
}
