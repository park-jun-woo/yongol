//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TM-51 단위 — 메뉴 렌더 항목 ≥1 + 호스트 레이아웃 전무 발화 / 레이아웃·defaultLayout·data-layout·메뉴0·hidden-subtree 침묵 검증

package stml_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM51_SitemapNoLayoutHost(t *testing.T) {
	doc := makeDoc(map[string]*openapi3.PathItem{
		"/home": getOp("GetHome", nil, nil),
	})
	pages := []stml.PageSpec{{Name: "home", FileName: "home.html"}}

	t.Run("fires when a menu renders but no layout hosts it", func(t *testing.T) {
		fs := makeFS(pages, doc)
		fs.Sitemap = menuSitemap()
		if got := countDiag(tm51SitemapNoLayoutHost(fs), "[TM-51]"); got != 1 {
			t.Errorf("expected 1 TM-51, got %d", got)
		}
	})

	t.Run("silent when a layout exists", func(t *testing.T) {
		fs := makeFS(pages, doc)
		fs.Sitemap = menuSitemap()
		fs.Layouts = []stml.LayoutSpec{{Name: "app", HasOutlet: true}}
		if got := countDiag(tm51SitemapNoLayoutHost(fs), "[TM-51]"); got != 0 {
			t.Errorf("expected 0 TM-51 with a layout, got %d", got)
		}
	})

	t.Run("silent when defaultLayout is declared (TM-12 owns it)", func(t *testing.T) {
		fs := makeFS(pages, doc)
		fs.Sitemap = menuSitemap()
		fs.Manifest.Frontend.DefaultLayout = "app"
		if got := countDiag(tm51SitemapNoLayoutHost(fs), "[TM-51]"); got != 0 {
			t.Errorf("expected 0 TM-51 with defaultLayout, got %d", got)
		}
	})

	t.Run("silent when a nav declares data-layout (TM-41 owns it)", func(t *testing.T) {
		fs := makeFS(pages, doc)
		fs.Sitemap = menuSitemap()
		fs.Sitemap.Navs[0].Layout = "app"
		if got := countDiag(tm51SitemapNoLayoutHost(fs), "[TM-51]"); got != 0 {
			t.Errorf("expected 0 TM-51 with nav data-layout, got %d", got)
		}
	})

	t.Run("silent when no menu entry renders (all data-menu=false)", func(t *testing.T) {
		fs := makeFS(pages, doc)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: "home", Label: "Home", Menu: false},
		}}}}
		if got := countDiag(tm51SitemapNoLayoutHost(fs), "[TM-51]"); got != 0 {
			t.Errorf("expected 0 TM-51 with no menu-rendered entry, got %d", got)
		}
	})

	// Regression guard (plan §검증 method): a data-menu="false" group parent
	// whose page child has its own menuBlockReason == "" must stay silent —
	// hidden-subtree propagation folds the child out of Roots, matching the
	// emitter. A per-node walk would have raised a spurious WARNING here.
	t.Run("silent for a hidden subtree with a per-node-renderable child", func(t *testing.T) {
		fs := makeFS([]stml.PageSpec{{Name: "detail", FileName: "detail.html"}}, doc)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Label: "Group", Menu: false, Children: []stml.SitemapNode{
				{Page: "detail", Label: "Detail", Menu: true},
			}},
		}}}}
		if got := countDiag(tm51SitemapNoLayoutHost(fs), "[TM-51]"); got != 0 {
			t.Errorf("expected 0 TM-51 for a hidden subtree, got %d", got)
		}
	})
}
