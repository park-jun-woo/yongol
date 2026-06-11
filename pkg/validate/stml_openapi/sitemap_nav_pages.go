//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what sitemapNavPages — 사이트맵 노드 트리의 모든 data-page 값을 재귀 수집 (data-entry 루트 편입용)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapNavPages collects every data-page value of a sitemap node tree,
// all depths, document order — buildPageGraph turns a data-entry block
// into reachability roots with it (every page of the block is a public
// entry point, DESIGN §4.10 root set).
func sitemapNavPages(nodes []stml.SitemapNode) []string {
	var out []string
	for _, n := range nodes {
		if n.Page != "" {
			out = append(out, n.Page)
		}
		out = append(out, sitemapNavPages(n.Children)...)
	}
	return out
}
