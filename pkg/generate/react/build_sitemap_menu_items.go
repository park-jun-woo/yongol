//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what buildSitemapMenuItems — 사이트맵 노드들을 메뉴 항목 모델로 변환 (MenuRenderable 판정 공유, 재귀)

package react

import (
	genstml "github.com/park-jun-woo/yongol/pkg/generate/react/stml"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/validate/stml_openapi"
)

// buildSitemapMenuItems converts sitemap nodes at the given 1-based menu
// depth into menu-item models. Admission is exactly the shared
// stml_openapi.MenuRenderable judgment (depth ≤ 2, no required route
// parameter, no data-menu="false") — TM-43 and this emitter must never
// disagree on what renders, and a data-menu="false" node hides its whole
// subtree because children are only built under admitted parents (the
// collectMenuRendered hidden-propagation semantics). Page items resolve
// their path through navLinkPath (the data-nav resolution) and gather the
// static route prefixes of their menu-hidden descendants for the
// ancestor-highlight className (DESIGN §4.4); icons translate kebab-case
// data-icon names to lucide-react PascalCase components. data-roles passes
// through as item.Roles (Phase005) — the renderer wraps the item in a role
// condition, and nesting alone realizes the subtree inheritance (children
// render inside the parent's conditional block, so ancestor conditions
// AND for free). A complete dynamic menu group (Phase007 — the shared
// stml_openapi.DynamicMenuGroup judgment) carries its fetch wiring and the
// item NavLink to attribute built by the page data-link emitter
// (genstml.LinkToAttr over the target's resolved route — the exact
// substitution a page <Link> gets); incomplete groups stay static, like
// validation (TM-48/TM-30) demands.
func buildSitemapMenuItems(nodes []stml.SitemapNode, depth int, routePatterns map[string]string) []sitemapMenuItem {
	var items []sitemapMenuItem
	for _, n := range nodes {
		if !stml_openapi.MenuRenderable(n, depth, sitemapNodePatterns(n, routePatterns)) {
			continue
		}
		item := sitemapMenuItem{Label: n.Label, Roles: n.Roles}
		if n.Icon != "" {
			item.Icon = kebabToPascal(n.Icon)
		}
		if stml_openapi.DynamicMenuGroup(n) {
			item.Fetch = n.Fetch
			item.Each = n.Each
			item.LabelField = n.LabelField
			item.ItemToAttr = genstml.LinkToAttr(stml.LinkRef{TargetPage: n.Link, Params: n.LinkParams, TargetPattern: routePatterns[n.Link]})
			item.ItemKey = sitemapDynamicItemKey(n.LinkParams)
		}
		switch {
		case n.Page != "":
			item.Kind = "page"
			item.To = navLinkPath(n.Page, routePatterns)
			var prefixes []string
			collectMenuActivePrefixes(n.Children, depth+1, routePatterns, &prefixes)
			item.Prefixes = dedupeStrings(prefixes)
		case n.Href != "":
			item.Kind = "external"
			item.Href = n.Href
		default:
			item.Kind = "group"
		}
		item.Children = buildSitemapMenuItems(n.Children, depth+1, routePatterns)
		items = append(items, item)
	}
	return items
}
