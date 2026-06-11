//ff:func feature=validate type=rule control=sequence topic=stml-openapi
//ff:what sitemapDiags — frontend/sitemap.html 존재 시의 사이트맵 규칙 일괄 실행 (TM-39~44, TM-46~48, TM-50, TM-51 + 동적 그룹 TM-01/07/30/31/32 확장)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// sitemapDiags runs every rule gated on an existing frontend/sitemap.html
// — Run's sitemap branch (plans/stml/sitemap Phase001~007): TM-39~42
// validate the tree itself, TM-43 the reachability BFS, TM-44 the menu's
// single source of truth, TM-46/47 the data-roles wiring, TM-50 the
// data-crumb-field declarations; dynamic menu groups (Phase007) get TM-48
// (data-entry block / structural completeness) plus the sitemap
// extensions of TM-01/07/30/31/32 (fetch op, each array field, label
// field, link target, link params). TM-51 is the inverse of TM-49: a
// sitemap that derives a menu but has no layout to host it (layouts/
// empty + no defaultLayout + no nav data-layout).
func sitemapDiags(fs *yongol.Fullstack, opMap map[string]operationEntry, raif map[string]map[string]map[string]bool) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, tm39SitemapPageNotFound(fs)...)
	diags = append(diags, tm40SitemapDuplicatePage(fs)...)
	diags = append(diags, tm41SitemapLayoutNotFound(fs)...)
	diags = append(diags, tm42SitemapIndexConflict(fs)...)
	diags = append(diags, tm43UnreachablePage(fs)...)
	diags = append(diags, tm44DataNavWithSitemap(fs)...)
	diags = append(diags, tm46SitemapRoleUnknown(fs)...)
	diags = append(diags, tm47RolesWiringMissing(fs)...)
	diags = append(diags, tm48SitemapDynamicGroup(fs)...)
	diags = append(diags, tm01SitemapGroupFetch(fs, opMap)...)
	diags = append(diags, tm07SitemapGroupEach(fs, opMap)...)
	diags = append(diags, tm30SitemapGroupLabelField(fs, raif)...)
	diags = append(diags, tm31SitemapGroupLink(fs)...)
	diags = append(diags, tm32SitemapGroupLinkParams(fs, raif)...)
	diags = append(diags, tm50CrumbField(fs, opMap)...)
	diags = append(diags, tm51SitemapNoLayoutHost(fs)...)
	return diags
}
