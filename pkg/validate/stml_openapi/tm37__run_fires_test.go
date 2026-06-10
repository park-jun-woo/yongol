//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-37 — 레이아웃 data-logout 미정의 op ERROR 발화, 정의된 POST op 시 해제 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM37_RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings":   getOp("ListBuildings", nil, nil),
		"/auth/logout": postOp("Logout", nil),
	})
	fs := makeFS([]stml.PageSpec{{
		Name:     "building-list",
		FileName: "building-list.html",
		Layout:   "app",
		Fetches:  []stml.FetchBlock{{OperationID: "ListBuildings"}},
	}}, doc)
	fs.Layouts = []stml.LayoutSpec{{
		Name:   "app",
		File:   "layouts/app.html",
		Logout: &stml.LogoutSpec{OperationID: "SignOut", Label: "로그아웃"},
	}}

	diags := Run(fs)
	if got := countDiag(diags, "[TM-37]"); got != 1 {
		t.Errorf("expected 1 TM-37 via Run (unknown op), got %d: %+v", got, diags)
	}

	fs.Layouts[0].Logout.OperationID = "Logout"
	diags = Run(fs)
	if got := countDiag(diags, "[TM-37]"); got != 0 {
		t.Errorf("expected 0 TM-37 via Run (declared POST op), got %d: %+v", got, diags)
	}
}
