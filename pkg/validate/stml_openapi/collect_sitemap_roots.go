//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectSitemapRoots — 사이트맵 nav 들을 순회하며 entry 루트 편입 + 메뉴 렌더 상태 기록

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectSitemapRoots feeds the sitemap's contribution to the root set
// into the graph: every page of a data-entry block (public entry points,
// DESIGN §4.10 root set — addEntryRoots) and the menu-rendered entries of
// every nav (collectMenuRendered, edge (a) — the menu hangs off the
// virtual root). names is the existing-page filter.
func collectSitemapRoots(sm *stml.SitemapSpec, names map[string]bool, pages []stml.PageSpec, g *pageGraph) {
	for _, nav := range sm.Navs {
		if nav.Entry {
			addEntryRoots(nav.Items, names, g)
		}
		collectMenuRendered(nav.Items, 1, "", pages, g)
	}
}
