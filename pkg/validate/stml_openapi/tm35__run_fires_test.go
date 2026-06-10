//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-35 — 인덱스 미선언 폴백 WARNING 발화, frontend.index 선언 시 해제 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM35RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{})
	pages := []stml.PageSpec{
		{Name: "forgot-password", FileName: "forgot-password.html"},
		{Name: "dashboard", FileName: "dashboard.html"},
	}

	fs := makeFS(pages, doc)
	diags := Run(fs)
	if got := countDiag(diags, "[TM-35]"); got != 1 {
		t.Errorf("expected 1 TM-35 via Run (undeclared index), got %d: %+v", got, diags)
	}

	declared := makeFS(pages, doc)
	declared.Manifest.Frontend.Index = "dashboard"
	diags = Run(declared)
	if got := countDiag(diags, "[TM-35]"); got != 0 {
		t.Errorf("expected 0 TM-35 via Run (index declared), got %d: %+v", got, diags)
	}
	if got := countDiag(diags, "[TM-34]"); got != 0 {
		t.Errorf("expected 0 TM-34 via Run (valid declaration), got %d: %+v", got, diags)
	}
}
