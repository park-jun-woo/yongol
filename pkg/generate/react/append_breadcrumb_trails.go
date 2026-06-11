//ff:func feature=gen-react type=generator control=iteration dimension=1
//ff:what appendBreadcrumbTrails — 사이트맵 노드 재귀 순회로 trail 누적 (깊이≥2 페이지만, 자신 crumb 은 href 없음)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// appendBreadcrumbTrails walks sitemap nodes depth-first, accumulating
// the ancestor crumb chain and appending one trail per page node
// (plans/stml/sitemap Phase004). A trail exists only for page nodes at
// depth ≥ 2 — a depth-1 page would render a single-crumb breadcrumb,
// which is noise (plan rule 4) — and only when the page resolves in
// routePatterns (an unresolved page is TM-39's finding, not an emission
// concern). The trail's own page is always label-only (current location);
// ancestor hrefs follow breadcrumbCrumbFor's MenuRenderable judgment.
// A data-crumb-field page marks its own crumb Dynamic (Phase006) — the
// runtime label slot; ancestor crumbs never (self label only, the plan's
// non-scope). Every node — group label, external link, page — contributes
// a crumb to its descendants' chains. depth is 1-based.
func appendBreadcrumbTrails(nodes []stml.SitemapNode, depth int, ancestors []breadcrumbCrumb, routePatterns map[string]string, trails *[]breadcrumbTrail) {
	for _, n := range nodes {
		crumb := breadcrumbCrumbFor(n, depth, routePatterns)
		if pattern, ok := routePatterns[n.Page]; ok && n.Page != "" && depth >= 2 {
			crumbs := make([]breadcrumbCrumb, 0, len(ancestors)+1)
			crumbs = append(crumbs, ancestors...)
			crumbs = append(crumbs, breadcrumbCrumb{Label: crumb.Label, Dynamic: n.CrumbField != ""})
			*trails = append(*trails, breadcrumbTrail{Page: n.Page, Pattern: pattern, Crumbs: crumbs})
		}
		appendBreadcrumbTrails(n.Children, depth+1, append(ancestors[:len(ancestors):len(ancestors)], crumb), routePatterns, trails)
	}
}
