//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what hasClaimsCapture — 클레임 캡처 탐지(이름 일치/페이지 부재) 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestHasClaimsCapture(t *testing.T) {
	pages := []stml.PageSpec{{FileName: "login.html", Actions: []stml.ActionBlock{{
		OperationID: "Login",
		Captures: []stml.CaptureBind{
			{RespField: "access_token", Sink: "auth.token"},
			{RespField: "role", Sink: "auth.claims.role"},
		},
	}}}}
	if !hasClaimsCapture(pages, "role") {
		t.Errorf("expected auth.claims.role capture to be found")
	}
	if hasClaimsCapture(pages, "user_role") {
		t.Errorf("auth.claims.role must not satisfy user_role")
	}
	if hasClaimsCapture(nil, "role") {
		t.Errorf("no pages: expected false")
	}
}
