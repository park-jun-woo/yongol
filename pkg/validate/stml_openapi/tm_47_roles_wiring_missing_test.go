//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-47 — role_field 미선언/auth.claims 캡처 부재/backend roles 빈 목록 각각 ERROR, 완전 배선·미사용 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/manifest"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM47RolesWiringMissing(t *testing.T) {
	roledSitemap := &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: "member-list", Label: "멤버", Menu: true, Roles: []string{"admin"}},
		}}},
	}
	capturing := stml.PageSpec{FileName: "login.html", Actions: []stml.ActionBlock{{
		OperationID: "Login",
		Captures:    []stml.CaptureBind{{RespField: "role", Sink: "auth.claims.role"}},
	}}}
	t.Run("fully wired is silent", func(t *testing.T) {
		fs := makeAuthFS([]stml.PageSpec{capturing}, nil, "cookie")
		fs.Manifest.Backend.Auth.Roles = []string{"admin", "member"}
		fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{RoleField: "role"}
		fs.Sitemap = roledSitemap
		if diags := tm47RolesWiringMissing(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("no data-roles use is silent even unwired", func(t *testing.T) {
		fs := makeAuthFS(nil, nil, "cookie")
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{{Page: "home", Label: "홈", Menu: true}}}}}
		if diags := tm47RolesWiringMissing(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("missing role_field fires", func(t *testing.T) {
		fs := makeAuthFS([]stml.PageSpec{capturing}, nil, "cookie")
		fs.Manifest.Backend.Auth.Roles = []string{"admin"}
		fs.Sitemap = roledSitemap
		diags := tm47RolesWiringMissing(fs)
		if countDiag(diags, "[TM-47]") != 1 || !strings.Contains(diags[0].Message, "role_field") {
			t.Errorf("expected 1 TM-47 about role_field, got %+v", diags)
		}
	})

	t.Run("missing claims capture fires", func(t *testing.T) {
		fs := makeAuthFS(nil, nil, "cookie")
		fs.Manifest.Backend.Auth.Roles = []string{"admin"}
		fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{RoleField: "role"}
		fs.Sitemap = roledSitemap
		diags := tm47RolesWiringMissing(fs)
		if countDiag(diags, "[TM-47]") != 1 || !strings.Contains(diags[0].Message, "auth.claims.role") {
			t.Errorf("expected 1 TM-47 about the missing capture, got %+v", diags)
		}
	})

	t.Run("capture of a different claim name does not satisfy", func(t *testing.T) {
		fs := makeAuthFS([]stml.PageSpec{capturing}, nil, "cookie")
		fs.Manifest.Backend.Auth.Roles = []string{"admin"}
		fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{RoleField: "user_role"}
		fs.Sitemap = roledSitemap
		if got := countDiag(tm47RolesWiringMissing(fs), "[TM-47]"); got != 1 {
			t.Errorf("expected 1 TM-47 (auth.claims.role ≠ auth.claims.user_role), got %d", got)
		}
	})

	t.Run("empty backend roles fires", func(t *testing.T) {
		fs := makeAuthFS([]stml.PageSpec{capturing}, nil, "cookie")
		fs.Manifest.Frontend.Auth = &manifest.FrontendAuth{RoleField: "role"}
		fs.Sitemap = roledSitemap
		diags := tm47RolesWiringMissing(fs)
		if countDiag(diags, "[TM-47]") != 1 || !strings.Contains(diags[0].Message, "backend.auth.roles") {
			t.Errorf("expected 1 TM-47 about empty roles, got %+v", diags)
		}
	})

	t.Run("everything missing fires role_field and roles errors", func(t *testing.T) {
		fs := makeAuthFS(nil, nil, "cookie")
		fs.Sitemap = roledSitemap
		if got := countDiag(tm47RolesWiringMissing(fs), "[TM-47]"); got != 2 {
			t.Errorf("expected 2 TM-47 (role_field + empty roles), got %d", got)
		}
	})
}
