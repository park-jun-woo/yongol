//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what MenuRenderable — 사이트맵 노드의 메뉴 렌더 판정 (깊이 ≤2 ∧ 필수 파라미터 없음 ∧ data-menu != "false") — Phase003 방출기와 공유하는 계약

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// MenuRenderable reports whether a sitemap node actually renders in the
// layout menu (plans/stml/sitemap Phase002, DESIGN §4.10): menu depth
// within the 2-level group → item limit AND the resolved route carries no
// required parameter segment (the TM-36 judgment over stml.RoutePaths)
// AND no explicit data-menu="false". This is the contract the Phase003
// menu emitter consumes as-is — validation and emission must never
// disagree on what renders. depth is 1-based; routePatterns is
// stml.RoutePaths of the node's page (nil for group labels and external
// links). The reason text lives in menuBlockReason so TM-43 can name it.
func MenuRenderable(node stml.SitemapNode, depth int, routePatterns []string) bool {
	return menuBlockReason(node, depth, routePatterns) == ""
}
