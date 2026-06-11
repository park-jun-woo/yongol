//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-07/08 사이트맵 확장 — 동적 메뉴 그룹의 data-each 가 응답 스키마에 없거나 배열이 아님 (tm0708Each 재사용)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm07SitemapGroupEach applies the TM-07/TM-08 judgment to sitemap
// dynamic menu groups (plans/stml/sitemap Phase007): the group's
// data-each must be an array field of the data-fetch operation's response
// schema — the exact tm0708Each check pages get, run over a synthesized
// EachBlock. Unknown operationIds stay silent (TM-01 owns them), as does
// a group still missing data-fetch or data-each (TM-48's finding).
func tm07SitemapGroupEach(fs *yongol.Fullstack, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, e := range sitemapDynamicGroupEntries(fs.Sitemap) {
		if e.Node.Fetch == "" || e.Node.Each == "" {
			continue
		}
		entry, ok := opMap[e.Node.Fetch]
		if !ok {
			continue
		}
		diags = append(diags, tm0708Each([]stml.EachBlock{{Field: e.Node.Each}}, e.Node.Fetch, fs.Sitemap.FileName, entry)...)
	}
	return diags
}
