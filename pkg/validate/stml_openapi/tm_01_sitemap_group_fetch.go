//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-01 사이트맵 확장 — 동적 메뉴 그룹의 data-fetch operationId 가 OpenAPI 에 없음 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm01SitemapGroupFetch applies the TM-01 judgment to sitemap dynamic
// menu groups (plans/stml/sitemap Phase007): the group's data-fetch must
// name an existing OpenAPI operationId, exactly like a page data-fetch —
// the layout emits the same api.<Op>() call a page would. A group without
// data-fetch is TM-48's structural finding.
func tm01SitemapGroupFetch(fs *yongol.Fullstack, opMap map[string]operationEntry) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, e := range sitemapDynamicGroupEntries(fs.Sitemap) {
		if e.Node.Fetch == "" {
			continue
		}
		if _, ok := opMap[e.Node.Fetch]; ok {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        fs.Sitemap.FileName,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-01] dynamic menu group data-fetch operationId %q at %s is not defined in OpenAPI", e.Node.Fetch, e.Path),
			Advice:      fmt.Sprintf("Add operationId %q to the OpenAPI spec, or fix the data-fetch value in sitemap.html", e.Node.Fetch),
			OperationID: e.Node.Fetch,
		})
	}
	return diags
}
