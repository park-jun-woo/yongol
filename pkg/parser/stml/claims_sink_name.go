//ff:func feature=stml-parse type=util control=sequence
//ff:what ClaimsSinkName — "auth.claims.<name>" sink 에서 claim 이름 추출 (식별자 검증 포함)
package stml

import "strings"

// ClaimsSinkName returns the <name> of an "auth.claims.<name>" capture
// sink and whether the sink is a well-formed claims sink: the name must be
// a non-empty identifier ([A-Za-z_][A-Za-z0-9_]*). It is the single
// judgment every consumer shares — the capture parser (whitelist), the
// capture emitter (setClaim), TM-24 (cookie-mode exception: claims come
// from the login response body, not the httpOnly cookie) and TM-47 (the
// role_field wiring check) — so they can never disagree on what counts as
// a claims sink (plans/stml/sitemap Phase005).
func ClaimsSinkName(sink string) (string, bool) {
	name, ok := strings.CutPrefix(sink, "auth.claims.")
	if !ok || !isIdentName(name) {
		return "", false
	}
	return name, true
}
