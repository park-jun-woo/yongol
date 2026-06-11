//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-32 사이트맵 확장 — 동적 메뉴 그룹들의 data-link-params 를 대상 라우트 패턴에 대해 검사 (그룹별 검사는 위임)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm32SitemapGroupLinkParams applies the TM-32 judgment to sitemap dynamic
// menu groups (plans/stml/sitemap Phase007): syntax, item.<Field> sources,
// segment existence and required-segment coverage against the target
// page's resolved route — the per-group check lives in
// tm32CheckSitemapGroup. A target page missing from the set stays silent
// (TM-31), as does an unresolvable item schema (TM-01/TM-07).
func tm32SitemapGroupLinkParams(fs *yongol.Fullstack, raif map[string]map[string]map[string]bool) []diagnostic.Diagnostic {
	patterns := map[string]string{}
	for _, p := range fs.STMLPages {
		if rp := stml.RoutePaths(p); len(rp) > 0 {
			patterns[p.Name] = rp[0]
		}
	}
	var diags []diagnostic.Diagnostic
	for _, e := range sitemapDynamicGroupEntries(fs.Sitemap) {
		diags = append(diags, tm32CheckSitemapGroup(e, fs.Sitemap.FileName, patterns, raif)...)
	}
	return diags
}
