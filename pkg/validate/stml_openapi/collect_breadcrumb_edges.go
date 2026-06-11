//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectBreadcrumbEdges — 사이트맵 nav 들을 순회하며 브레드크럼 상행 간선(자식→MenuRenderable 조상) 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectBreadcrumbEdges feeds the breadcrumb up-edges of DESIGN §4.10
// edge (d) into the graph (plans/stml/sitemap Phase004): the generated
// <Breadcrumb> renders, on every sitemap-listed page of depth ≥ 2, a link
// to each ancestor the shared MenuRenderable judgment admits — exactly
// the crumbs the Phase004 emitter gives an href, so validation and
// emission never disagree on what the user can click. names is the
// existing-page filter; the per-node walk lives in addBreadcrumbEdges.
func collectBreadcrumbEdges(sm *stml.SitemapSpec, names map[string]bool, pages []stml.PageSpec, g *pageGraph) {
	for _, nav := range sm.Navs {
		addBreadcrumbEdges(nav.Items, 1, nil, names, pages, g)
	}
}
