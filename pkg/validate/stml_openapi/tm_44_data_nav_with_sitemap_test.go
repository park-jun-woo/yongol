//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-44 — data-nav 잔존 레이아웃마다 ERROR + 마이그레이션 안내, data-nav 없는 레이아웃 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM44DataNavWithSitemap(t *testing.T) {
	t.Run("each layout with data-nav fires", func(t *testing.T) {
		fs := makeFS(nil, nil)
		fs.Layouts = []stml.LayoutSpec{
			{Name: "app", File: "layouts/app.html", NavItems: []stml.NavItem{{Path: "dashboard", Label: "대시보드"}}},
			{Name: "admin", File: "layouts/admin.html", NavItems: []stml.NavItem{{Path: "/admin", Label: "관리"}}},
			{Name: "bare", File: "layouts/bare.html"},
		}
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html"}

		diags := tm44DataNavWithSitemap(fs)
		if got := countDiag(diags, "[TM-44]"); got != 2 {
			t.Fatalf("expected 2 TM-44 (one per data-nav layout), got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, "메뉴는 sitemap.html 로 이동") {
			t.Errorf("Message must carry the migration guidance, got %q", diags[0].Message)
		}
	})

	t.Run("layouts without data-nav are silent", func(t *testing.T) {
		fs := makeFS(nil, nil)
		fs.Layouts = []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html"}
		if diags := tm44DataNavWithSitemap(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
