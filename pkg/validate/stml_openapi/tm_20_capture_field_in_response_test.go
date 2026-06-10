//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM20CaptureFieldInResponse — 구문 위반·sink 위반·미존재 필드·정상 캡처 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM20CaptureFieldInResponse(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/login": postOpWithResp("Login", map[string]*openapi3.SchemaRef{
			"access_token":  stringProp(),
			"refresh_token": stringProp(),
		}),
	})
	opMap := buildOperationMethodMap(doc)

	// Valid captures → no diagnostics.
	ok := tm20CaptureFieldInResponse(stml.ActionBlock{
		OperationID: "Login",
		CaptureRaw:  "access_token -> auth.token, refresh_token -> auth.refresh",
	}, "login.html", opMap)
	if len(ok) != 0 {
		t.Errorf("valid: expected 0 diagnostics, got %+v", ok)
	}

	// Typo response field → 1 ERROR.
	typo := tm20CaptureFieldInResponse(stml.ActionBlock{
		OperationID: "Login",
		CaptureRaw:  "acces_token -> auth.token",
	}, "login.html", opMap)
	if countDiag(typo, "[TM-20]") != 1 {
		t.Errorf("typo field: expected 1 TM-20, got %+v", typo)
	}

	// Syntax violation (missing arrow) → 1 ERROR.
	syntax := tm20CaptureFieldInResponse(stml.ActionBlock{
		OperationID: "Login",
		CaptureRaw:  "access_token auth.token",
	}, "login.html", opMap)
	if countDiag(syntax, "[TM-20]") != 1 {
		t.Errorf("syntax: expected 1 TM-20, got %+v", syntax)
	}

	// Disallowed sink → 1 ERROR.
	sink := tm20CaptureFieldInResponse(stml.ActionBlock{
		OperationID: "Login",
		CaptureRaw:  "access_token -> session.token",
	}, "login.html", opMap)
	if countDiag(sink, "[TM-20]") != 1 {
		t.Errorf("sink: expected 1 TM-20, got %+v", sink)
	}

	// No data-capture → no diagnostics.
	if d := tm20CaptureFieldInResponse(stml.ActionBlock{OperationID: "Login"}, "login.html", opMap); len(d) != 0 {
		t.Errorf("no capture: expected 0 diagnostics, got %+v", d)
	}

	// Unknown operationId → silent (TM-02 reports it).
	if d := tm20CaptureFieldInResponse(stml.ActionBlock{
		OperationID: "Nope",
		CaptureRaw:  "access_token -> auth.token",
	}, "login.html", opMap); len(d) != 0 {
		t.Errorf("unknown op: expected 0 diagnostics, got %+v", d)
	}
}
