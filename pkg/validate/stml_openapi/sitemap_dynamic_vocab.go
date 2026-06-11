//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what sitemapDynamicVocab — 노드에 선언된 동적 그룹 속성명 목록 (선언 순서 고정) — "어휘 사용" 판정 공유

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapDynamicVocab returns the dynamic-group attribute names a sitemap
// node declares (plans/stml/sitemap Phase007), in the fixed vocabulary
// order. A non-empty result marks the node as a dynamic-group declaration
// — complete or not — which is what TM-48 and the per-attribute rules key
// off; the names feed the diagnostics verbatim.
func sitemapDynamicVocab(node stml.SitemapNode) []string {
	var attrs []string
	if node.Fetch != "" {
		attrs = append(attrs, "data-fetch")
	}
	if node.Each != "" {
		attrs = append(attrs, "data-each")
	}
	if node.Link != "" {
		attrs = append(attrs, "data-link")
	}
	if node.LinkParamsRaw != "" {
		attrs = append(attrs, "data-link-params")
	}
	if node.LabelField != "" {
		attrs = append(attrs, "data-label-field")
	}
	return attrs
}
