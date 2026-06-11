//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-46 — 미선언 role 값 ERROR(메뉴 숨김 ≠ 보안 명시) / 선언 role 침묵 / roles 빈 목록 위임 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM46SitemapRoleUnknown(t *testing.T) {
	sitemap := func(roles ...string) *stml.SitemapSpec {
		return &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
				{Page: "member-list", Label: "멤버", Menu: true, Roles: roles},
			}}},
		}
	}

	t.Run("unknown role fires one ERROR naming the spot", func(t *testing.T) {
		fs := makeAuthFS(nil, nil, "cookie")
		fs.Manifest.Backend.Auth.Roles = []string{"member", "admin"}
		fs.Sitemap = sitemap("adimn")
		diags := tm46SitemapRoleUnknown(fs)
		if countDiag(diags, "[TM-46]") != 1 {
			t.Fatalf("expected 1 TM-46, got %+v", diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "멤버") {
			t.Errorf("Message should carry the tree position, got %q", diags[0].Message)
		}
		if !strings.Contains(diags[0].Message, "menu hiding is not security") || !strings.Contains(diags[0].Message, "Rego") {
			t.Errorf("Message should state menu hiding ≠ access blocking (Rego), got %q", diags[0].Message)
		}
		if !strings.Contains(diags[0].Advice, "member, admin") {
			t.Errorf("Advice should list the declared roles, got %q", diags[0].Advice)
		}
	})

	t.Run("each unknown value of a multi-role list fires", func(t *testing.T) {
		fs := makeAuthFS(nil, nil, "cookie")
		fs.Manifest.Backend.Auth.Roles = []string{"member"}
		fs.Sitemap = sitemap("admin", "member", "manager")
		if got := countDiag(tm46SitemapRoleUnknown(fs), "[TM-46]"); got != 2 {
			t.Errorf("expected 2 TM-46, got %d", got)
		}
	})

	t.Run("declared roles are silent", func(t *testing.T) {
		fs := makeAuthFS(nil, nil, "cookie")
		fs.Manifest.Backend.Auth.Roles = []string{"member", "admin"}
		fs.Sitemap = sitemap("admin", "member")
		if diags := tm46SitemapRoleUnknown(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})

	t.Run("empty backend.auth.roles defers to TM-47", func(t *testing.T) {
		fs := makeAuthFS(nil, nil, "cookie")
		fs.Sitemap = sitemap("admin")
		if diags := tm46SitemapRoleUnknown(fs); len(diags) != 0 {
			t.Errorf("expected silence (TM-47's finding), got %+v", diags)
		}
	})
}
