//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO10_LayoutLogoutConsumed — 레이아웃 data-logout 만이 소비하는 op 은 XMO-10 미발화 (오발화 차단 고정)

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO10_LayoutLogoutConsumed(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "building-list",
		FileName: "building-list.html",
		Layout:   "app",
		Fetches:  []stml.FetchBlock{{OperationID: "ListBuildings"}},
	}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings":   getOp("ListBuildings", nil, nil),
		"/auth/logout": postOp("Logout", nil),
	})
	fs := makeFS(pages, doc)
	fs.Layouts = []stml.LayoutSpec{{
		Name:   "app",
		File:   "layouts/app.html",
		Logout: &stml.LogoutSpec{OperationID: "Logout", Label: "로그아웃"},
	}}

	diags := Run(fs)
	if hasDiag(diags, "[XMO-10]") {
		t.Errorf("op consumed by layout data-logout must not trigger XMO-10, got %v", diags)
	}

	// Without the logout declaration the same op is unconsumed — the
	// rule still guards real gaps.
	fs.Layouts[0].Logout = nil
	diags = Run(fs)
	if got := countDiag(diags, "[XMO-10]"); got != 1 {
		t.Errorf("expected 1 XMO-10 without the layout consumer, got %d: %+v", got, diags)
	}
}
