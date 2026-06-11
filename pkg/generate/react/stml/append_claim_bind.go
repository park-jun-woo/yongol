//ff:func feature=stml-gen type=util control=sequence
//ff:what appendClaimBind — 바인딩이 auth.claims.* sink 일 때만 claims 목록에 추가

package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// appendClaimBind appends c to claims only when its sink is a well-formed
// auth.claims.* sink (stmlparser.ClaimsSinkName — the shared judgment).
func appendClaimBind(claims []stmlparser.CaptureBind, c stmlparser.CaptureBind) []stmlparser.CaptureBind {
	if _, ok := stmlparser.ClaimsSinkName(c.Sink); !ok {
		return claims
	}
	return append(claims, c)
}
