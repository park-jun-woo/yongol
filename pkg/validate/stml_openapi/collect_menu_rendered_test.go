//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what collectMenuRendered — 등재 기록 / 렌더→루트 편입 / 비렌더 사유 기록 / data-menu="false" 서브트리 차단 검증

package stml_openapi

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCollectMenuRendered(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "building-list", FileName: "building-list.html", Route: "/buildings"},
		{Name: "building-detail", FileName: "building-detail.html", Route: "/buildings/:BuildingID"},
		{Name: "settings", FileName: "settings.html", Route: "/settings"},
	}
	newGraph := func() *pageGraph {
		return &pageGraph{Roots: map[string]bool{}, Edges: map[string][]string{}, InSitemap: map[string]bool{}, MenuBlocked: map[string]string{}}
	}

	t.Run("rendered page joins the roots, blocked one records its reason", func(t *testing.T) {
		g := newGraph()
		nodes := []stml.SitemapNode{
			{Label: "건물 관리", Menu: true, Children: []stml.SitemapNode{
				{Page: "building-list", Menu: true, Children: []stml.SitemapNode{
					{Page: "building-detail", Menu: true},
				}},
			}},
		}
		collectMenuRendered(nodes, 1, "", pages, g)
		if !g.Roots["building-list"] {
			t.Error("building-list (depth 2, no params) should be menu-rendered and a root")
		}
		if g.Roots["building-detail"] {
			t.Error("building-detail must not be a root — listing is a node, not an edge")
		}
		if !g.InSitemap["building-detail"] {
			t.Error("building-detail should be recorded as listed")
		}
		if reason := g.MenuBlocked["building-detail"]; !strings.Contains(reason, "depth 3") {
			t.Errorf("MenuBlocked reason = %q, want the depth reason", reason)
		}
	})

	t.Run("data-menu=false hides the whole subtree", func(t *testing.T) {
		g := newGraph()
		nodes := []stml.SitemapNode{
			{Label: "숨김 그룹", Menu: false, Children: []stml.SitemapNode{
				{Page: "settings", Menu: true},
			}},
		}
		collectMenuRendered(nodes, 1, "", pages, g)
		if g.Roots["settings"] {
			t.Error("settings under a hidden subtree must not be menu-rendered")
		}
		if reason := g.MenuBlocked["settings"]; !strings.Contains(reason, `data-menu="false"`) {
			t.Errorf("MenuBlocked reason = %q, want the hidden-subtree reason", reason)
		}
	})
}
