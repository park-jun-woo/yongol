//ff:func feature=gen-gogin type=util control=selection
//ff:what isResourceIDZero — @auth Inputs 의 ResourceID 표현식이 정적으로 zero 인지 판정

package ssac

import "strings"

// isResourceIDZero reports whether the ResourceID expression pulled
// from a `@auth` sequence's Inputs map is statically zero — i.e. the
// handler is a creation form and no resource exists yet. Detection is
// by inspection of the SSaC AST expression; runtime zeros (e.g. a
// variable that happens to resolve to 0) are out of scope. Matched
// literals: empty string, `0`, `""`, `nil`, `null` (case-insensitive),
// with surrounding whitespace ignored.
//
// The caller (buildAuth) combines this with a presence check: a
// missing ResourceID key also counts as zero. See Phase005 (BUG-033).
func isResourceIDZero(expr string) bool {
	s := strings.TrimSpace(expr)
	if s == "" {
		return true
	}
	switch strings.ToLower(s) {
	case "0", `""`, "''", "nil", "null":
		return true
	}
	return false
}
