//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what recordMenuPage — 사이트맵 페이지 노드 하나의 등재/루트 편입/비렌더 사유를 pageGraph 에 기록

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// recordMenuPage records one sitemap node's page into the graph for
// collectMenuRendered: listed (InSitemap), menu-rendered (reason "" →
// Roots, edge (a) of DESIGN §4.10), otherwise the first recorded
// menu-block reason (MenuBlocked — TM-43's cause text). Group labels and
// external links carry no page and are skipped.
func recordMenuPage(node stml.SitemapNode, reason string, g *pageGraph) {
	if node.Page == "" {
		return
	}
	g.InSitemap[node.Page] = true
	if reason == "" {
		g.Roots[node.Page] = true
		return
	}
	if _, seen := g.MenuBlocked[node.Page]; !seen {
		g.MenuBlocked[node.Page] = reason
	}
}
