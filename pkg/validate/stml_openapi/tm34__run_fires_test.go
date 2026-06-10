//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-34 — frontend.index 오타 페이지명이 ERROR 로 발화함을 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM34RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{})
	fs := makeFS([]stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
	}, doc)
	fs.Manifest.Frontend.Index = "dashbord" // typo

	diags := Run(fs)
	if got := countDiag(diags, "[TM-34]"); got != 1 {
		t.Errorf("expected 1 TM-34 via Run, got %d: %+v", got, diags)
	}
	// index is declared (even if typo'd) → the fallback warning stays out.
	if got := countDiag(diags, "[TM-35]"); got != 0 {
		t.Errorf("expected 0 TM-35 via Run, got %d: %+v", got, diags)
	}
}
