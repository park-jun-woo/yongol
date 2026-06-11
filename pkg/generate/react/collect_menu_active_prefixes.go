//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what collectMenuActivePrefixes — 메뉴 비렌더 자손 페이지의 정적 라우트 prefix 수집 (가장 가까운 메뉴 렌더 조상 하이라이트)

package react

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/validate/stml_openapi"
)

// collectMenuActivePrefixes walks a rendered menu item's subtree and
// gathers the static route prefixes of descendant pages that do NOT
// render in the menu themselves — those pages highlight their nearest
// menu-rendered ancestor (DESIGN §4.4 activeMenu auto-derivation). A
// child the shared stml_openapi.MenuRenderable judgment admits is skipped
// whole: it renders, so its own subtree highlights it, not this item.
// Useless prefixes ("" and "/") are dropped. depth is 1-based, the level
// of the nodes passed in.
func collectMenuActivePrefixes(nodes []stml.SitemapNode, depth int, routePatterns map[string]string, out *[]string) {
	for _, n := range nodes {
		if stml_openapi.MenuRenderable(n, depth, sitemapNodePatterns(n, routePatterns)) {
			continue
		}
		// routePatterns[""] is "" for group labels/external links, so the
		// prefix check alone filters them out alongside useless prefixes.
		if p := staticRoutePrefix(routePatterns[n.Page]); p != "" && p != "/" {
			*out = append(*out, p)
		}
		collectMenuActivePrefixes(n.Children, depth+1, routePatterns, out)
	}
}
