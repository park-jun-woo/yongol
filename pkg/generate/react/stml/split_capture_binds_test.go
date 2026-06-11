//ff:func feature=stml-gen type=test control=sequence
//ff:what splitCaptureBinds — token/refresh 필드 추출·claims 분류·미지 sink 무시·빈 입력 검증
package stml

import (
	"reflect"
	"testing"

	stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestSplitCaptureBinds(t *testing.T) {
	captures := []stmlparser.CaptureBind{
		{RespField: "access_token", Sink: "auth.token"},
		{RespField: "refresh_token", Sink: "auth.refresh"},
		{RespField: "role", Sink: "auth.claims.role"},
		{RespField: "tier", Sink: "auth.claims.tier"},
		{RespField: "bogus", Sink: "session.user"}, // outside the whitelist → no store write
	}
	tokenField, refreshField, claims := splitCaptureBinds(captures)
	if tokenField != "access_token" {
		t.Errorf("tokenField = %q, want %q", tokenField, "access_token")
	}
	if refreshField != "refresh_token" {
		t.Errorf("refreshField = %q, want %q", refreshField, "refresh_token")
	}
	wantClaims := []stmlparser.CaptureBind{
		{RespField: "role", Sink: "auth.claims.role"},
		{RespField: "tier", Sink: "auth.claims.tier"},
	}
	if !reflect.DeepEqual(claims, wantClaims) {
		t.Errorf("claims = %+v, want %+v", claims, wantClaims)
	}

	// no captures → zero values across the board
	tokenField, refreshField, claims = splitCaptureBinds(nil)
	if tokenField != "" || refreshField != "" || claims != nil {
		t.Errorf("empty input: got (%q, %q, %+v), want (\"\", \"\", nil)", tokenField, refreshField, claims)
	}
}
