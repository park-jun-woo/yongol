//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectMenuRendered — 사이트맵 트리를 걸으며 등재/메뉴 렌더(루트 편입)/비렌더 사유를 pageGraph 에 기록

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectMenuRendered walks one sitemap nav's nodes depth-first, recording
// every data-page node into the graph via recordMenuPage: listed, whether
// it actually renders in the menu (MenuRenderable → folded into Roots,
// since menu edges emanate from the virtual root — DESIGN §4.10 edge (a)),
// and otherwise the human-readable reason TM-43 names. A data-menu="false"
// node hides its whole subtree: descendants inherit hiddenReason instead
// of their own judgment, because items under an unrendered entry cannot
// render either. A rendered complete dynamic menu group (plans/stml/sitemap
// Phase007) folds its data-link target into the roots too — every fetched
// item is a menu NavLink to that page, the same edge-(b) movement a page
// data-link supplies, hanging off the virtual root like edge (a). depth is
// 1-based.
func collectMenuRendered(nodes []stml.SitemapNode, depth int, hiddenReason string, pages []stml.PageSpec, g *pageGraph) {
	for _, n := range nodes {
		reason := hiddenReason
		if reason == "" {
			reason = menuBlockReason(n, depth, nodeRoutePatterns(n, pages))
		}
		recordMenuPage(n, reason, g)
		if reason == "" && DynamicMenuGroup(n) && findPageByName(pages, n.Link) != nil {
			g.Roots[n.Link] = true
		}
		childHidden := hiddenReason
		if childHidden == "" && !n.Menu {
			childHidden = `inside a subtree hidden by data-menu="false"`
		}
		collectMenuRendered(n.Children, depth+1, childHidden, pages, g)
	}
}
