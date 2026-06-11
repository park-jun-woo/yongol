//ff:func feature=gen-react type=util control=sequence
//ff:what sitemapNodePatterns — 사이트맵 노드의 라우트 패턴 해석 (그룹/외부 링크/미해석 페이지는 nil) — MenuRenderable 인자 공급

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapNodePatterns resolves a sitemap node's route patterns for the
// shared stml_openapi.MenuRenderable judgment — the react-side twin of the
// validator's nodeRoutePatterns: the page's resolved pattern from
// routePatterns (navRoutePatterns — stml.RoutePaths first pattern), nil
// for group labels, external links and unresolved pages (the latter is
// TM-39's finding, not a menu judgment).
func sitemapNodePatterns(node stml.SitemapNode, routePatterns map[string]string) []string {
	if node.Page == "" {
		return nil
	}
	p, ok := routePatterns[node.Page]
	if !ok {
		return nil
	}
	return []string{p}
}
