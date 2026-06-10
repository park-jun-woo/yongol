//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-38 — cookie 모드 값 없는 data-logout WARNING 발화, op 선언 시 해제 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM38_RunFires(t *testing.T) {
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
	fs.Manifest.Backend.Auth = &manifest.Auth{Type: "jwt", Mode: "cookie"}
	fs.Layouts = []stml.LayoutSpec{{
		Name:   "app",
		File:   "layouts/app.html",
		Logout: &stml.LogoutSpec{Label: "로그아웃"},
	}}

	diags := Run(fs)
	if got := countDiag(diags, "[TM-38]"); got != 1 {
		t.Errorf("expected 1 TM-38 via Run (cookie + valueless), got %d: %+v", got, diags)
	}

	fs.Layouts[0].Logout.OperationID = "Logout"
	diags = Run(fs)
	if got := countDiag(diags, "[TM-38]"); got != 0 {
		t.Errorf("expected 0 TM-38 via Run (server op declared), got %d: %+v", got, diags)
	}
}
