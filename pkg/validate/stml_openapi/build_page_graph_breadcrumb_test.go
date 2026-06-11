//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-openapi
//ff:what buildPageGraph 간선 (d) — 브레드크럼 상행 간선이 도달성을 확장 / redirect 중복 간선 dedupe / TM-43 침묵 검증

package stml_openapi

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestBuildPageGraphBreadcrumbEdges(t *testing.T) {
	// The Phase004 scenario where edge (d) alone reaches a page: the whole
	// tools subtree is menu-hidden (data-menu="false" propagation — none of
	// it is a Root), dashboard data-links straight to tool-detail, and
	// tool-detail's breadcrumb links up to tool-list. Without edge (d)
	// tool-list would be a false orphan.
	pages := []stml.PageSpec{
		{Name: "dashboard", FileName: "dashboard.html", Children: []stml.ChildNode{
			{Kind: "link", Link: &stml.LinkRef{TargetPage: "tool-detail"}},
		}},
		{Name: "tools", FileName: "tools.html"},
		{Name: "tool-list", FileName: "tool-list.html"},
		// the redirect duplicates the breadcrumb edge target — appendEdgeOnce dedupes
		{Name: "tool-detail", FileName: "tool-detail.html", Route: "/tools/:ToolID", Actions: []stml.ActionBlock{
			{OperationID: "DeleteTool", Redirect: "tool-list"},
		}},
	}
	fs := makeFS(pages, nil)
	fs.Sitemap = &stml.SitemapSpec{
		FileName: "sitemap.html",
		Navs: []stml.SitemapNav{{Items: []stml.SitemapNode{
			{Page: "dashboard", Label: "대시보드", Index: true, Menu: true},
			{Page: "tools", Label: "도구", Menu: false, Children: []stml.SitemapNode{
				{Page: "tool-list", Label: "도구 목록", Menu: true, Children: []stml.SitemapNode{
					{Page: "tool-detail", Label: "도구 상세", Menu: true},
				}},
			}},
		}}},
	}

	g := buildPageGraph(fs)
	// redirect edge first (page-edge loop), breadcrumb edge deduped after it
	if want := []string{"tool-list"}; !reflect.DeepEqual(g.Edges["tool-detail"], want) {
		t.Errorf("Edges[tool-detail] = %v, want %v (redirect + breadcrumb deduped)", g.Edges["tool-detail"], want)
	}
	if g.Roots["tool-list"] {
		t.Error("tool-list must not be a root (hidden subtree) — only edge (d) reaches it")
	}
	reached := reachablePages(g)
	for _, name := range []string{"dashboard", "tool-detail", "tool-list"} {
		if !reached[name] {
			t.Errorf("%q should be reachable (dashboard → link → detail → breadcrumb up-edge)", name)
		}
	}
	if reached["tools"] {
		t.Error("tools (data-menu=false, no crumb href) must stay unreachable")
	}
	diags := tm43UnreachablePage(fs)
	if got := countDiag(diags, "[TM-43]"); got != 1 {
		t.Fatalf("expected exactly 1 TM-43 (tools), got %d: %+v", got, diags)
	}
}
