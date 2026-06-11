//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-39 — sitemap data-page 가 실존 STML 페이지가 아님 + data-page 와 <a href> 동시 보유 + 그룹 li 의 data-crumb-field (ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm39SitemapPageNotFound checks every sitemap data-page value against the
// existing STML page names (plans/stml/sitemap Phase001). An entry holding
// both data-page and an <a href> external link is rejected here too — the
// two are mutually exclusive vehicles (internal page vs external URL), and
// the parser deliberately preserves the contradiction for this rule to name.
// data-crumb-field on a page-less <li> (group label / external link) is
// the same family of structural misplacement (Phase006 — the dynamic
// crumb label is read from the page's fetch, so only a page item can
// carry it); TM-50 validates the field itself.
func tm39SitemapPageNotFound(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	names := make(map[string]struct{}, len(fs.STMLPages))
	for _, p := range fs.STMLPages {
		names[p.Name] = struct{}{}
	}

	var diags []diagnostic.Diagnostic
	for _, e := range collectSitemapEntries(fs.Sitemap) {
		if e.Node.Page == "" {
			diags = append(diags, tm39CrumbFieldMisplaced(e, fs.Sitemap.FileName)...)
			continue
		}
		if e.Node.Href != "" {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.Sitemap.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-39] sitemap entry at %s declares both data-page=%q and an <a href=%q> external link — they are mutually exclusive", e.Path, e.Node.Page, e.Node.Href),
				Advice:  "Keep one: data-page for an internal STML page, or an <a href> child for an external link",
			})
		}
		if _, ok := names[e.Node.Page]; !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.Sitemap.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-39] sitemap data-page %q at %s does not name any STML page", e.Node.Page, e.Path),
				Advice:  fmt.Sprintf("Use the page's STML filename without .html, or create frontend/%s.html", e.Node.Page),
			})
		}
	}
	return diags
}
