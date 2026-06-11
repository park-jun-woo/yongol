//ff:func feature=gen-react type=test control=iteration dimension=1
//ff:what buildRoutes 3단 사슬 — 페이지 data-layout > sitemap 블록 data-layout > defaultLayout 검증

package react

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildRoutes_SitemapLayoutChain(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "a", FileName: "a.html", Layout: "own"}, // page wins over sitemap
		{Name: "b", FileName: "b.html"},                // sitemap wins over default
		{Name: "c", FileName: "c.html"},                // falls to defaultLayout
	}
	sitemapLayouts := map[string]string{"a": "admin", "b": "admin"}

	routes := buildRoutes(pages, "app", nil, sitemapLayouts)
	got := map[string]string{}
	for _, r := range routes {
		got[r.Path] = r.Layout
	}
	if got["/a"] != "own" {
		t.Errorf(`page data-layout must win: /a layout = %q, want "own"`, got["/a"])
	}
	if got["/b"] != "admin" {
		t.Errorf(`sitemap block layout must beat defaultLayout: /b layout = %q, want "admin"`, got["/b"])
	}
	if got["/c"] != "app" {
		t.Errorf(`defaultLayout fallback broken: /c layout = %q, want "app"`, got["/c"])
	}
}
