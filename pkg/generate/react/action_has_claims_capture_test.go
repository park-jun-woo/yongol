//ff:func feature=gen-react type=test control=sequence
//ff:what actionHasClaimsCapture — claims sink 존재/token·refresh 만/비식별자 sink/캡처 없음 판정 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestActionHasClaimsCapture(t *testing.T) {
	// at least one auth.claims.* capture → true, even mixed with tokens
	withClaims := stml.ActionBlock{Captures: []stml.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
		{RespField: "role", Sink: "auth.claims.role"},
	}}
	if !actionHasClaimsCapture(withClaims) {
		t.Error("action with auth.claims.role capture: got false, want true")
	}

	// only token/refresh sinks → false
	tokensOnly := stml.ActionBlock{Captures: []stml.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
		{RespField: "refresh_token", Sink: "auth.refresh"},
	}}
	if actionHasClaimsCapture(tokensOnly) {
		t.Error("token/refresh-only action: got true, want false")
	}

	// malformed claims sink (non-identifier name) is not a claims capture —
	// the gate shares stml.ClaimsSinkName with the parser's whitelist
	malformed := stml.ActionBlock{Captures: []stml.CaptureBind{
		{RespField: "role", Sink: "auth.claims.ro-le"},
	}}
	if actionHasClaimsCapture(malformed) {
		t.Error("malformed claims sink: got true, want false")
	}

	// no captures at all → false
	if actionHasClaimsCapture(stml.ActionBlock{}) {
		t.Error("action without captures: got true, want false")
	}
}
