//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectSitemapEntries — SitemapSpec 의 모든 노드를 위치 경로와 함께 문서 순서로 평탄화

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// collectSitemapEntries flattens every node of the sitemap (all nav blocks,
// all depths) into document order, each paired with its position path.
func collectSitemapEntries(sm *stml.SitemapSpec) []sitemapEntry {
	var entries []sitemapEntry
	for i, nav := range sm.Navs {
		appendSitemapEntries(nav.Items, fmt.Sprintf("nav[%d]", i), &entries)
	}
	return entries
}
