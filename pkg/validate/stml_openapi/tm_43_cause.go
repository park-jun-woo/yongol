//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what tm43Cause — 미도달 페이지의 원인 분류 문구 조립 (메뉴 비렌더 등재 / 완전 미등재 × 들어오는 간선 유무)

package stml_openapi

import "fmt"

// tm43Cause builds the cause clause of a TM-43 message: the sitemap axis
// (listed but not menu-rendered, with the recorded reason, vs not listed
// at all — listing is a node, not an edge) and the incoming-edge axis (no
// data-link/data-redirect/breadcrumb edge at all — since Phase004 a
// reachable descendant's breadcrumb up-link also counts — vs only
// unreachable pages pointing at it). Naming both axes tells the user
// which side to fix.
func tm43Cause(g *pageGraph, name string, hasIncoming bool) string {
	reason := g.MenuBlocked[name]
	if reason == "" {
		reason = "not menu-rendered"
	}
	sitemapPart := "not in the sitemap"
	if g.InSitemap[name] {
		sitemapPart = fmt.Sprintf("in sitemap but not menu-rendered (%s)", reason)
	}
	incomingPart := "no data-link/data-redirect/breadcrumb edge points to it"
	if hasIncoming {
		incomingPart = "only unreachable pages link to it"
	}
	return sitemapPart + ", and " + incomingPart
}
