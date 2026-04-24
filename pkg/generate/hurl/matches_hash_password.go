//ff:func feature=gen-hurl type=util control=sequence
//ff:what matchesHashPassword — @call Model 필드가 auth.HashPassword 대상인지 판정

package hurl

import (
	"strings"
)

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
