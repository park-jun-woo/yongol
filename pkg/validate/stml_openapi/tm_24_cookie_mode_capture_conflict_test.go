//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM24CookieModeCaptureConflict — cookie 모드의 캡처·frontend.auth 모순 WARNING + auth.claims.*·role_field 전용 블록 예외 검증

package stml_openapi

import (
	"strings"
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

	// cookie + auth.claims.* capture → exempt (the claim comes from the
	// login response body, not the httpOnly cookie — Phase005).
	claimsCapturing := stml.PageSpec{FileName: "login.html", Actions: []stml.ActionBlock{{
		OperationID: "Login",
		CaptureRaw:  "role -> auth.claims.role",
		Captures:    []stml.CaptureBind{{RespField: "role", Sink: "auth.claims.role"}},
	}}}
	fs = makeAuthFS([]stml.PageSpec{claimsCapturing}, doc, "cookie")
	if d := tm24CookieModeCaptureConflict(fs); len(d) != 0 {
		t.Errorf("claims capture: expected 0 diagnostics, got %+v", d)
	}

	// cookie + mixed captures (token + claims) → only the token capture warns.
	mixedCapturing := stml.PageSpec{FileName: "login.html", Actions: []stml.ActionBlock{{
		OperationID: "Login",
		CaptureRaw:  "access_token -> auth.token, role -> auth.claims.role",
		Captures: []stml.CaptureBind{
			{RespField: "access_token", Sink: "auth.token"},
			{RespField: "role", Sink: "auth.claims.role"},
		},
	}}}
	fs = makeAuthFS([]stml.PageSpec{mixedCapturing}, doc, "cookie")
	if countDiag(tm24CookieModeCaptureConflict(fs), "[TM-24]") != 1 {
		t.Errorf("mixed captures: expected 1 TM-24, got %+v", tm24CookieModeCaptureConflict(fs))
	}

	// cookie + role_field-only frontend.auth block → exempt (the same
	// RoleFieldOnly predicate XON-60 consumes).
	fs = makeAuthFS([]stml.PageSpec{plain}, doc, "cookie")
	fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{RoleField: "role"}
	if d := tm24CookieModeCaptureConflict(fs); len(d) != 0 {
		t.Errorf("role_field-only block: expected 0 diagnostics, got %+v", d)
	}

	// cookie + mixed block (token_field + role_field) → 1 WARNING with the
	// token-keys-only advice.
	fs = makeAuthFS([]stml.PageSpec{plain}, doc, "cookie")
	fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{TokenField: "access_token", RoleField: "role"}
	mixed := tm24CookieModeCaptureConflict(fs)
	if countDiag(mixed, "[TM-24]") != 1 {
		t.Fatalf("mixed block: expected 1 TM-24, got %+v", mixed)
	}
	if !strings.Contains(mixed[0].Advice, "token-related keys") {
		t.Errorf("mixed block advice should ask to remove only the token-related keys, got %q", mixed[0].Advice)
	}

	// bearer mode → rule skipped even with captures.
	fs = makeAuthFS([]stml.PageSpec{capturing}, doc, "bearer")
	if d := tm24CookieModeCaptureConflict(fs); len(d) != 0 {
		t.Errorf("bearer: expected 0 diagnostics, got %+v", d)
	}
}
