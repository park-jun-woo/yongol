//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-41 — 미실존 data-layout ERROR / 실존·미선언 레이아웃 침묵 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM41SitemapLayoutNotFound(t *testing.T) {
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}

	t.Run("nonexistent layout fires", func(t *testing.T) {
		fs := makeFS(nil, nil)
		fs.Layouts = layouts
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs:     []stml.SitemapNav{{Layout: "ghost"}},
		}
		diags := tm41SitemapLayoutNotFound(fs)
		if got := countDiag(diags, "[TM-41]"); got != 1 {
			t.Fatalf("expected 1 TM-41, got %d: %+v", got, diags)
		}
		if diags[0].Level != diagnostic.LevelError {
			t.Errorf("Level = %v, want LevelError", diags[0].Level)
		}
		if !strings.Contains(diags[0].Message, `"ghost"`) {
			t.Errorf("Message should name the layout, got %q", diags[0].Message)
		}
	})

	t.Run("existing layout and empty layout are silent", func(t *testing.T) {
		fs := makeFS(nil, nil)
		fs.Layouts = layouts
		fs.Sitemap = &stml.SitemapSpec{
			FileName: "sitemap.html",
			Navs:     []stml.SitemapNav{{Layout: "app"}, {Layout: ""}},
		}
		if diags := tm41SitemapLayoutNotFound(fs); len(diags) != 0 {
			t.Errorf("expected silence, got %+v", diags)
		}
	})
}
