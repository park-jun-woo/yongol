//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what nodeRoutePatterns — 사이트맵 노드의 페이지를 해석해 stml.RoutePaths 반환 (그룹/외부 링크/미실존 페이지는 nil)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// nodeRoutePatterns resolves the route patterns of a sitemap node's page
// for the MenuRenderable judgment: stml.RoutePaths of the named page, nil
// when the node is a group label / external link or the page does not
// exist (the latter is TM-39's finding, not a menu judgment).
func nodeRoutePatterns(node stml.SitemapNode, pages []stml.PageSpec) []string {
	if node.Page == "" {
		return nil
	}
	target := findPageByName(pages, node.Page)
	if target == nil {
		return nil
	}
	return stml.RoutePaths(*target)
}
