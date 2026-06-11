//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what sitemapDynamicGroupEntries — 동적 그룹 어휘를 하나라도 가진 사이트맵 노드들을 위치 경로와 함께 평탄화

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// sitemapDynamicGroupEntries flattens the sitemap to the nodes declaring
// any dynamic-group vocabulary (plans/stml/sitemap Phase007), each paired
// with its position path — the shared walk of the dynamic-group rules
// (TM-01/07/30/31/32 sitemap extensions and TM-48's completeness check).
func sitemapDynamicGroupEntries(sm *stml.SitemapSpec) []sitemapEntry {
	var out []sitemapEntry
	for _, e := range collectSitemapEntries(sm) {
		if len(sitemapDynamicVocab(e.Node)) > 0 {
			out = append(out, e)
		}
	}
	return out
}
