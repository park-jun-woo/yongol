//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what actionHasClaimsCapture — 액션의 캡처 바인딩에 auth.claims.* sink 존재 판정

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// actionHasClaimsCapture reports whether the action declares at least one
// auth.claims.* capture binding (stml.ClaimsSinkName — the shared
// judgment, so the gate cannot drift from the parser's whitelist).
func actionHasClaimsCapture(a stml.ActionBlock) bool {
	for _, c := range a.Captures {
		if _, ok := stml.ClaimsSinkName(c.Sink); ok {
			return true
		}
	}
	return false
}
