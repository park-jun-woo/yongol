//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what isSignupShape — SSaC ServiceFunc 가 auth.HashPassword 호출을 포함하는지 판정

package hurl

import (
	"strings"

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

// matchesHashPassword checks for the `auth.HashPassword` target in an
// @call sequence's Model field. parseCallExprInputs keeps the full
// "pkg.Func" form for @call (splitPackagePrefix is only applied to
// @get/@post/@put/@delete), so a direct suffix match on ".HashPassword"
// with a leading "auth" package segment is correct.
func matchesHashPassword(model string) bool {
	if model == "" {
		return false
	}
	// Exact form seen in yongol: "auth.HashPassword"
	if model == "auth.HashPassword" {
		return true
	}
	// Defensive: tolerate whitespace or trailing tokens (should not occur
	// after parser, but keeps the check robust to future syntax tweaks).
	m := strings.TrimSpace(model)
	return m == "auth.HashPassword" || strings.HasSuffix(m, ".auth.HashPassword")
}

// hasUserCreatePost checks for a companion `@post <Model>.Create` sequence
// that wires a PasswordHash-like column. Used only to emit a WARNING
// when signup shape is suspicious (HashPassword without a Create post).
func hasUserCreatePost(fn *ssac.ServiceFunc) bool {
	if fn == nil {
		return false
	}
	for _, seq := range fn.Sequences {
		if seq.Type != ssac.SeqPost {
			continue
		}
		if !strings.HasSuffix(seq.Model, ".Create") {
			continue
		}
		for k := range seq.Inputs {
			if strings.EqualFold(k, "PasswordHash") {
				return true
			}
		}
	}
	return false
}
