//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what menuBlockReason — 사이트맵 노드가 메뉴에 렌더되지 않는 이유 반환 (렌더되면 "") — MenuRenderable 의 판정 본체

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// menuBlockReason returns why a sitemap node would not render in the
// layout menu, "" when it renders. The three judgments of DESIGN §4.7
// (plans/stml/sitemap Phase002), in precedence order: an explicit
// data-menu="false", a menu depth beyond the 2-level (group → item)
// render limit, and a resolved route carrying a required parameter
// segment (the TM-36 judgment via firstRequiredSegment — a static menu
// link has no value to fill it). depth is 1-based (1 = top-level entry);
// routePatterns is stml.RoutePaths of the node's page (nil for groups,
// external links and unresolved pages).
func menuBlockReason(node stml.SitemapNode, depth int, routePatterns []string) string {
	if !node.Menu {
		return `hidden by data-menu="false"`
	}
	if depth > 2 {
		return fmt.Sprintf("menu depth %d exceeds the 2-level (group → item) render limit", depth)
	}
	seg := ""
	if len(routePatterns) > 0 {
		seg = firstRequiredSegment(routePatterns[0])
	}
	if seg != "" {
		return fmt.Sprintf("route %q has required param %s", routePatterns[0], seg)
	}
	return ""
}
