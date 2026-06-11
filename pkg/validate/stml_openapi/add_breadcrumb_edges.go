//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what addBreadcrumbEdges — 사이트맵 노드 재귀 순회로 자식 페이지→MenuRenderable 조상 페이지 간선 누적

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// addBreadcrumbEdges walks sitemap nodes depth-first, carrying the chain
// of linkable ancestors — existing pages the raw MenuRenderable judgment
// admits at their own depth, the exact href condition of the Phase004
// breadcrumb emitter (a crumb link is a static route, so it works even
// inside a data-menu="false" subtree; no hidden-propagation here, unlike
// the edge (a) root folding). Every existing page node of depth ≥ 2 gets
// one edge per linkable ancestor (DESIGN §4.10 edge (d), added through
// addBreadcrumbPageEdges); depth-1 pages have no ancestors and render no
// breadcrumb at all. depth is 1-based.
func addBreadcrumbEdges(nodes []stml.SitemapNode, depth int, ancestors []string, names map[string]bool, pages []stml.PageSpec, g *pageGraph) {
	for _, n := range nodes {
		if n.Page != "" && names[n.Page] {
			addBreadcrumbPageEdges(n.Page, ancestors, g)
		}
		next := ancestors
		if n.Page != "" && names[n.Page] && MenuRenderable(n, depth, nodeRoutePatterns(n, pages)) {
			next = append(ancestors[:len(ancestors):len(ancestors)], n.Page)
		}
		addBreadcrumbEdges(n.Children, depth+1, next, names, pages, g)
	}
}
