//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what addBreadcrumbPageEdges — 한 페이지의 브레드크럼 상행 간선들을 그래프에 추가 (자기 자신 제외, 중복 억제)

package stml_openapi

// addBreadcrumbPageEdges adds one page's breadcrumb up-edges: page →
// every linkable ancestor of its trail (DESIGN §4.10 edge (d), Phase004).
// The self guard is defensive — a page cannot be its own sitemap ancestor
// under TM-40's canonical rule — and appendEdgeOnce dedupes against the
// data-link/data-redirect edges collected before the sitemap walk.
func addBreadcrumbPageEdges(page string, ancestors []string, g *pageGraph) {
	for _, a := range ancestors {
		if a == page {
			continue
		}
		appendEdgeOnce(g, page, a)
	}
}
