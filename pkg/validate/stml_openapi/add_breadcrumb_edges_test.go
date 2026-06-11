//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what addBreadcrumbEdges — data-menu=false 숨김 서브트리 내부 조상도 raw MenuRenderable 이면 간선 대상 (href 는 정적 라우트) 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestAddBreadcrumbEdges(t *testing.T) {
	// tools hides its whole subtree from the menu (data-menu="false"), so
	// none of these pages is a Root. The breadcrumb still links tool-list
	// from tool-detail — a crumb href is a static route and works without
	// a menu entry (raw MenuRenderable, no hidden propagation) — while
	// tools itself (data-menu="false" → not renderable) gets no edge.
	pages := []stml.PageSpec{
		{Name: "tools", FileName: "tools.html"},
		{Name: "tool-list", FileName: "tool-list.html"},
		{Name: "tool-detail", FileName: "tool-detail.html", Route: "/tools/:ToolID"},
	}
	names := map[string]bool{"tools": true, "tool-list": true, "tool-detail": true}
	nodes := []stml.SitemapNode{
		{Page: "tools", Label: "도구", Menu: false, Children: []stml.SitemapNode{
			{Page: "tool-list", Label: "도구 목록", Menu: true, Children: []stml.SitemapNode{
				{Page: "tool-detail", Label: "도구 상세", Menu: true},
			}},
		}},
	}

	g := &pageGraph{Edges: map[string][]string{}}
	addBreadcrumbEdges(nodes, 1, nil, names, pages, g)

	if want := []string{"tool-list"}; !reflect.DeepEqual(g.Edges["tool-detail"], want) {
		t.Errorf("Edges[tool-detail] = %v, want %v (hidden-subtree ancestor stays linkable, data-menu=false ancestor does not)", g.Edges["tool-detail"], want)
	}
	if got := g.Edges["tool-list"]; got != nil {
		t.Errorf("Edges[tool-list] = %v, want none (its only ancestor is not menu-renderable)", got)
	}
}
