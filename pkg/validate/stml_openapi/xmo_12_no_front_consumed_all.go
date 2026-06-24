//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what XMO-12 (도메인) — 전체 도메인 OpenAPI 의 no-front 태그 op 가 전체 도메인 STML/레이아웃/사이트맵/컴포넌트에서 실제 소비 중 (WARNING)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xmo12NoFrontConsumedAll is the domain-mode XMO-12. It builds the consumption
// set across every domain (each domain's pages, layouts, and sitemap, plus the
// shared component scan filtered against the union of operationIds), then warns
// for any no-front-tagged operation in any domain's document that is actually
// consumed. operationIds are globally unique under XDO-90, so merging across
// domains never conflates two operations.
func xmo12NoFrontConsumedAll(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if !frontendEnabled(fs) {
		return nil
	}

	docs := fs.AllOpenAPIDocs()
	ops := make(map[string]struct{})
	for _, doc := range docs {
		for id := range collectOpIDs(doc) {
			ops[id] = struct{}{}
		}
	}

	consumed := make(map[string]struct{})
	for _, name := range fs.DomainNames() {
		sub := collectConsumedOps(fs.DomainSTMLPages[name], fs.DomainLayouts[name], fs.DomainSitemaps[name], fs.SpecsDir, ops)
		for id := range sub {
			consumed[id] = struct{}{}
		}
	}

	var diags []diagnostic.Diagnostic
	for _, name := range fs.DomainNames() {
		doc := fs.DomainOpenAPIDocs[name]
		if doc == nil || doc.Paths == nil {
			continue
		}
		for _, item := range doc.Paths.Map() {
			diags = append(diags, xmo12ItemDiags(item, consumed)...)
		}
	}
	return diags
}
