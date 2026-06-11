//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-31 사이트맵 확장 — 동적 메뉴 그룹의 data-link 대상 페이지명이 STML 페이지 집합에 없음 (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm31SitemapGroupLink applies the TM-31 judgment to sitemap dynamic menu
// groups (plans/stml/sitemap Phase007): the group's data-link must name an
// existing STML page (filename without .html) — every fetched item becomes
// a NavLink to that page, so a typo'd target would emit a whole group of
// links into the void. A group without data-link is TM-48's finding.
func tm31SitemapGroupLink(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, e := range sitemapDynamicGroupEntries(fs.Sitemap) {
		if e.Node.Link == "" || findPageByName(fs.STMLPages, e.Node.Link) != nil {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fs.Sitemap.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-31] dynamic menu group data-link target %q at %s does not name any STML page", e.Node.Link, e.Path),
			Advice:  "Use the target page's STML filename without .html (a page-name reference, not a path)",
		})
	}
	return diags
}
