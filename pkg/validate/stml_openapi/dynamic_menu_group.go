//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what DynamicMenuGroup — 사이트맵 노드의 동적 메뉴 그룹 완전성 판정 (fetch·each·link·label-field 전부 선언) — 방출기와 공유하는 계약

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// DynamicMenuGroup reports whether a sitemap node is a complete dynamic
// menu group (plans/stml/sitemap Phase007, DESIGN §4.11 (a)): data-fetch,
// data-each, data-link and data-label-field all declared on the node's
// nested <ul>. data-link-params is required only by the target route's
// segments (TM-32's judgment). This is the contract the menu emitter and
// the reachability graph consume as-is — an incomplete group (TM-48 /
// TM-30's findings) emits nothing and contributes no edge, so validation
// and emission never disagree on what renders.
func DynamicMenuGroup(node stml.SitemapNode) bool {
	return node.Fetch != "" && node.Each != "" && node.Link != "" && node.LabelField != ""
}
