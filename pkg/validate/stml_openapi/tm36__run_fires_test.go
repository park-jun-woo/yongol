//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-36 — 레이아웃 data-nav 오타 ERROR 발화, 정상 페이지명 시 해제 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM36_RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings": getOp("ListBuildings", nil, nil),
	})
	fs := makeFS([]stml.PageSpec{{
		Name:     "building-list",
		FileName: "building-list.html",
		Layout:   "app",
		Fetches:  []stml.FetchBlock{{OperationID: "ListBuildings"}},
	}}, doc)
	fs.Layouts = []stml.LayoutSpec{{
		Name:     "app",
		File:     "layouts/app.html",
		NavItems: []stml.NavItem{{Path: "building-detial", Label: "건물"}},
	}}

	diags := Run(fs)
	if got := countDiag(diags, "[TM-36]"); got != 1 {
		t.Errorf("expected 1 TM-36 via Run (typo page name), got %d: %+v", got, diags)
	}

	fs.Layouts[0].NavItems[0].Path = "building-list"
	diags = Run(fs)
	if got := countDiag(diags, "[TM-36]"); got != 0 {
		t.Errorf("expected 0 TM-36 via Run (valid page name), got %d: %+v", got, diags)
	}
}
