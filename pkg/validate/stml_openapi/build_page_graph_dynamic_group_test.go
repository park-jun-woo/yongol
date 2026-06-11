//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what buildPageGraph — 동적 메뉴 그룹 data-link 대상의 루트 편입(렌더 그룹)·비렌더/불완전/미존재 시 미편입 검증

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildPageGraphDynamicGroup(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html"},
		{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID"},
	}
	group := stml.SitemapNode{Label: "내 건물", Menu: true, Fetch: "ListMyBuildings", Each: "items", Link: "building-detail", LinkParamsRaw: "item.building_id -> BuildingID", LabelField: "building_name"}
	graphFor := func(nodes []stml.SitemapNode) *pageGraph {
		fs := makeFS(pages, nil)
		fs.Sitemap = &stml.SitemapSpec{FileName: "sitemap.html", Navs: []stml.SitemapNav{{Items: nodes}}}
		return buildPageGraph(fs)
	}

	t.Run("a rendered complete group folds its link target into the roots", func(t *testing.T) {
		g := graphFor([]stml.SitemapNode{{Page: "dashboard", Label: "대시보드", Index: true, Menu: true}, group})
		if !g.Roots["building-detail"] {
			t.Errorf("Roots = %+v, want building-detail folded in (edge (b) of the dynamic items)", g.Roots)
		}
	})

	t.Run("a menu-hidden group contributes no edge", func(t *testing.T) {
		hidden := group
		hidden.Menu = false
		g := graphFor([]stml.SitemapNode{{Page: "dashboard", Label: "대시보드", Index: true, Menu: true}, hidden})
		if g.Roots["building-detail"] {
			t.Errorf("a data-menu=\"false\" group must not reach its target")
		}
	})

	t.Run("an incomplete group contributes no edge", func(t *testing.T) {
		incomplete := group
		incomplete.LabelField = ""
		g := graphFor([]stml.SitemapNode{{Page: "dashboard", Label: "대시보드", Index: true, Menu: true}, incomplete})
		if g.Roots["building-detail"] {
			t.Errorf("an incomplete group (TM-48/TM-30's finding) must not reach its target")
		}
	})

	t.Run("a nonexistent link target is not folded in", func(t *testing.T) {
		ghost := group
		ghost.Link = "nope-detail"
		g := graphFor([]stml.SitemapNode{{Page: "dashboard", Label: "대시보드", Index: true, Menu: true}, ghost})
		if g.Roots["nope-detail"] {
			t.Errorf("a nonexistent target (TM-31's finding) must not enter the roots")
		}
	})
}
