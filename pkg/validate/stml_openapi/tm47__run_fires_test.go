//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what Run 경유 TM-47 — data-roles 사용·배선 미완 ERROR 발화, 완전 배선 시 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM47_RunFires(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/members": getOp("ListMembers", nil, nil),
		"/login":   postOpWithResp("Login", map[string]*openapi3.SchemaRef{"role": stringProp()}),
	})
	pages := []stml.PageSpec{
		{Name: "member-list", FileName: "member-list.html", Fetches: []stml.FetchBlock{{OperationID: "ListMembers"}}},
		{Name: "login", FileName: "login.html", Actions: []stml.ActionBlock{{
			OperationID: "Login",
			Redirect:    "/",
		}}},
	}
	fs := makeAuthFS(pages, doc, "cookie")
	fs.Manifest.Backend.Auth.Roles = []string{"member", "admin"}
	fs.Sitemap = &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{
			Entry: true,
			Items: []stml.SitemapNode{
				{Page: "login", Label: "로그인", Index: true, Menu: true},
				{Page: "member-list", Label: "멤버", Menu: true, Roles: []string{"admin"}},
			},
		}},
	}

	// data-roles used but no role_field and no claims capture → TM-47.
	if got := countDiag(Run(fs), "[TM-47]"); got != 1 {
		t.Errorf("expected 1 TM-47 via Run, got %d: %+v", got, Run(fs))
	}

	// Full wiring (role_field + auth.claims.role capture + roles) → silent.
	fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{RoleField: "role"}
	fs.STMLPages[1].Actions[0].CaptureRaw = "role -> auth.claims.role"
	fs.STMLPages[1].Actions[0].Captures = []stml.CaptureBind{{RespField: "role", Sink: "auth.claims.role"}}
	diags := Run(fs)
	if got := countDiag(diags, "[TM-47]"); got != 0 {
		t.Errorf("expected 0 TM-47 when fully wired, got %d: %+v", got, diags)
	}
	// The cookie-mode claims capture must not draw TM-24 either (the
	// Phase005 exemption), nor TM-20 (the field exists in the response).
	if got := countDiag(diags, "[TM-24]"); got != 0 {
		t.Errorf("expected 0 TM-24 for the cookie-mode claims capture, got %d: %+v", got, diags)
	}
	if got := countDiag(diags, "[TM-20]"); got != 0 {
		t.Errorf("expected 0 TM-20 for the existing response field, got %d: %+v", got, diags)
	}
}
