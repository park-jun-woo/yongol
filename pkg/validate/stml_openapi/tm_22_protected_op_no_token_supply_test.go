//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM22ProtectedOpNoTokenSupply — bearer 보호 op 호출 + 캡처 전무 ERROR·정상·스킵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM22ProtectedOpNoTokenSupply(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/login": postOpWithResp("Login", map[string]*openapi3.SchemaRef{"access_token": stringProp()}),
		"/me":    securedGetOp("GetMe"),
	})
	opMap := buildOperationMethodMap(doc)

	plainLogin := stml.PageSpec{FileName: "login.html", Actions: []stml.ActionBlock{{OperationID: "Login"}}}
	mePage := stml.PageSpec{FileName: "me.html", Fetches: []stml.FetchBlock{{OperationID: "GetMe"}}}

	// bearer + protected consumer + no auth.token capture → 1 ERROR.
	fs := makeAuthFS([]stml.PageSpec{plainLogin, mePage}, doc, "bearer")
	got := tm22ProtectedOpNoTokenSupply(fs, opMap)
	if countDiag(got, "[TM-22]") != 1 || got[0].Level != diagnostic.LevelError {
		t.Errorf("fire: expected 1 TM-22 ERROR, got %+v", got)
	}

	// Token capture present → no diagnostics.
	capturing := stml.PageSpec{FileName: "login.html", Actions: []stml.ActionBlock{{
		OperationID: "Login",
		CaptureRaw:  "access_token -> auth.token",
		Captures:    []stml.CaptureBind{{RespField: "access_token", Sink: "auth.token"}},
	}}}
	fs = makeAuthFS([]stml.PageSpec{capturing, mePage}, doc, "bearer")
	if d := tm22ProtectedOpNoTokenSupply(fs, opMap); len(d) != 0 {
		t.Errorf("captured: expected 0 diagnostics, got %+v", d)
	}

	// No page calls a protected op → no diagnostics.
	fs = makeAuthFS([]stml.PageSpec{plainLogin}, doc, "bearer")
	if d := tm22ProtectedOpNoTokenSupply(fs, opMap); len(d) != 0 {
		t.Errorf("no protected page: expected 0 diagnostics, got %+v", d)
	}

	// cookie mode → rule skipped.
	fs = makeAuthFS([]stml.PageSpec{plainLogin, mePage}, doc, "cookie")
	if d := tm22ProtectedOpNoTokenSupply(fs, opMap); len(d) != 0 {
		t.Errorf("cookie: expected 0 diagnostics, got %+v", d)
	}
}
