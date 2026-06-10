//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestXMO12_LayoutLogoutConsumed — no-front 태그 op 을 레이아웃 data-logout 이 소비 중이면 XMO-12 WARNING

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestXMO12_LayoutLogoutConsumed(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "building-list",
		FileName: "building-list.html",
		Layout:   "app",
		Fetches:  []stml.FetchBlock{{OperationID: "ListBuildings"}},
	}}
	noFrontLogout := &openapi3.PathItem{Post: &openapi3.Operation{OperationID: "Logout", Tags: []string{"no-front"}}}
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/buildings":   getOp("ListBuildings", nil, nil),
		"/auth/logout": noFrontLogout,
	})
	fs := makeFS(pages, doc)
	fs.Layouts = []stml.LayoutSpec{{
		Name:   "app",
		File:   "layouts/app.html",
		Logout: &stml.LogoutSpec{OperationID: "Logout", Label: "로그아웃"},
	}}

	diags := Run(fs)
	if got := countDiag(diags, "[XMO-12]"); got != 1 {
		t.Errorf("expected 1 XMO-12 (stale no-front consumed by layout), got %d: %+v", got, diags)
	}

	// Dropping the layout consumer clears the stale-tag warning.
	fs.Layouts[0].Logout = nil
	diags = Run(fs)
	if got := countDiag(diags, "[XMO-12]"); got != 0 {
		t.Errorf("expected 0 XMO-12 without the layout consumer, got %d: %+v", got, diags)
	}
}
