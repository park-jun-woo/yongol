//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM24CookieModeCaptureConflict — cookie 모드의 캡처·frontend.auth 모순 WARNING 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM24CookieModeCaptureConflict(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/login": postOpWithResp("Login", map[string]*openapi3.SchemaRef{"access_token": stringProp()}),
	})
	capturing := stml.PageSpec{FileName: "login.html", Actions: []stml.ActionBlock{{
		OperationID: "Login",
		CaptureRaw:  "access_token -> auth.token",
		Captures:    []stml.CaptureBind{{RespField: "access_token", Sink: "auth.token"}},
	}}}
	plain := stml.PageSpec{FileName: "login.html", Actions: []stml.ActionBlock{{OperationID: "Login"}}}

	// cookie + capture → 1 WARNING.
	fs := makeAuthFS([]stml.PageSpec{capturing}, doc, "cookie")
	if countDiag(tm24CookieModeCaptureConflict(fs), "[TM-24]") != 1 {
		t.Errorf("capture: expected 1 TM-24, got %+v", tm24CookieModeCaptureConflict(fs))
	}

	// cookie (defaulted mode) + frontend.auth declared → 1 WARNING.
	fs = makeAuthFS([]stml.PageSpec{plain}, doc, "")
	fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{TokenField: "access_token"}
	if countDiag(tm24CookieModeCaptureConflict(fs), "[TM-24]") != 1 {
		t.Errorf("frontend.auth: expected 1 TM-24, got %+v", tm24CookieModeCaptureConflict(fs))
	}

	// cookie + capture + frontend.auth → 2 WARNINGs.
	fs = makeAuthFS([]stml.PageSpec{capturing}, doc, "cookie")
	fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{TokenField: "access_token"}
	if countDiag(tm24CookieModeCaptureConflict(fs), "[TM-24]") != 2 {
		t.Errorf("both: expected 2 TM-24, got %+v", tm24CookieModeCaptureConflict(fs))
	}

	// cookie, clean declarations → no diagnostics.
	fs = makeAuthFS([]stml.PageSpec{plain}, doc, "cookie")
	if d := tm24CookieModeCaptureConflict(fs); len(d) != 0 {
		t.Errorf("clean cookie: expected 0 diagnostics, got %+v", d)
	}

	// bearer mode → rule skipped even with captures.
	fs = makeAuthFS([]stml.PageSpec{capturing}, doc, "bearer")
	if d := tm24CookieModeCaptureConflict(fs); len(d) != 0 {
		t.Errorf("bearer: expected 0 diagnostics, got %+v", d)
	}
}
