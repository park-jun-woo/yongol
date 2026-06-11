//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what sitemapDeclaresIndex — 사이트맵에 data-index 노드가 존재하는지 판정 (TM-35 인덱스 선언 수단 확장)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapDeclaresIndex reports whether the sitemap marks any entry with
// data-index — the third index-declaration vehicle next to data-route="/"
// and manifest.frontend.index (plans/stml/sitemap Phase001). nil sitemap =
// no declaration.
func sitemapDeclaresIndex(sm *stml.SitemapSpec) bool {
	if sm == nil {
		return false
	}
	for _, e := range collectSitemapEntries(sm) {
		if e.Node.Index {
			return true
		}
	}
	return false
}
