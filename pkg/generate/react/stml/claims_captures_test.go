//ff:func feature=stml-gen type=test control=sequence
//ff:what claimsCaptures — auth.claims.* sink 만 순서대로 필터/비식별자 제외/없으면 nil 검증
package stml

import (
	"reflect"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestClaimsCaptures(t *testing.T) {
	binds := []stmlparser.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
		{RespField: "role", Sink: "auth.claims.role"},
		{RespField: "refresh_token", Sink: "auth.refresh"},
		{RespField: "tier", Sink: "auth.claims.tier"},
		{RespField: "bad", Sink: "auth.claims.ro-le"}, // non-identifier name → excluded
	}
	got := claimsCaptures(binds)
	want := []stmlparser.CaptureBind{
		{RespField: "role", Sink: "auth.claims.role"},
		{RespField: "tier", Sink: "auth.claims.tier"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}

	// token/refresh-only bindings → nil
	tokensOnly := []stmlparser.CaptureBind{{RespField: "access_token", Sink: "auth.token"}}
	if got := claimsCaptures(tokensOnly); got != nil {
		t.Errorf("tokens-only: got %+v, want nil", got)
	}
}
