//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM21CaptureSinkUnused — bearer 캡처 0건·소비처 0건·정상·cookie 스킵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM21CaptureSinkUnused(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/login": postOpWithResp("Login", map[string]*openapi3.SchemaRef{"access_token": stringProp()}),
		"/me":    securedGetOp("GetMe"),
	})
	opMap := buildOperationMethodMap(doc)

	loginAction := stml.ActionBlock{
		OperationID: "Login",
		CaptureRaw:  "access_token -> auth.token",
		Captures:    []stml.CaptureBind{{RespField: "access_token", Sink: "auth.token"}},
	}
	loginPage := stml.PageSpec{FileName: "login.html", Actions: []stml.ActionBlock{loginAction}}
	mePage := stml.PageSpec{FileName: "me.html", Fetches: []stml.FetchBlock{{OperationID: "GetMe"}}}

	// bearer + token capture + protected consumer → no diagnostics.
	fs := makeAuthFS([]stml.PageSpec{loginPage, mePage}, doc, "bearer")
	if d := tm21CaptureSinkUnused(fs, opMap); len(d) != 0 {
		t.Errorf("healthy bearer: expected 0 diagnostics, got %+v", d)
	}

	// bearer + no captures at all → 1 WARNING (no auth.token capture).
	plainLogin := stml.PageSpec{FileName: "login.html", Actions: []stml.ActionBlock{{OperationID: "Login"}}}
	fs = makeAuthFS([]stml.PageSpec{plainLogin, mePage}, doc, "bearer")
	if countDiag(tm21CaptureSinkUnused(fs, opMap), "[TM-21]") != 1 {
		t.Errorf("no captures: expected 1 TM-21, got %+v", tm21CaptureSinkUnused(fs, opMap))
	}

	// bearer + capture but no page calls a protected op → 1 WARNING.
	fs = makeAuthFS([]stml.PageSpec{loginPage}, doc, "bearer")
	if countDiag(tm21CaptureSinkUnused(fs, opMap), "[TM-21]") != 1 {
		t.Errorf("no consumer: expected 1 TM-21, got %+v", tm21CaptureSinkUnused(fs, opMap))
	}

	// cookie mode → rule skipped entirely.
	fs = makeAuthFS([]stml.PageSpec{plainLogin, mePage}, doc, "cookie")
	if d := tm21CaptureSinkUnused(fs, opMap); len(d) != 0 {
		t.Errorf("cookie: expected 0 diagnostics, got %+v", d)
	}

	// no backend.auth → rule skipped.
	fs = makeFS([]stml.PageSpec{plainLogin, mePage}, doc)
	if d := tm21CaptureSinkUnused(fs, opMap); len(d) != 0 {
		t.Errorf("no auth: expected 0 diagnostics, got %+v", d)
	}
}
