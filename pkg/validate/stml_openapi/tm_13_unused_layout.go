//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-13 — layouts/에 정의된 레이아웃이 어떤 페이지·defaultLayout·sitemap 블록에서도 사용되지 않음 (WARNING)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm13UnusedLayout detects layouts that are neither referenced by any
// page's data-layout attribute, nor set as the manifest's defaultLayout,
// nor assigned by a sitemap <nav data-sitemap data-layout> block
// (plans/stml/sitemap Phase003 — the menu emitter and the route builder
// use that assignment, so omitting it here would be a false "unused"
// against a layout the generator actually emits a menu for; validation
// and emission must not drift). sitemap may be nil.
func tm13UnusedLayout(pages []stml.PageSpec, layouts []stml.LayoutSpec, defaultLayout string, sitemap *stml.SitemapSpec) []diagnostic.Diagnostic {
	used := make(map[string]struct{})
	if defaultLayout != "" {
		used[defaultLayout] = struct{}{}
	}
	for _, page := range pages {
		if page.Layout != "" {
			used[page.Layout] = struct{}{}
		}
	}
	var sitemapNavs []stml.SitemapNav
	if sitemap != nil {
		sitemapNavs = sitemap.Navs
	}
	for _, nav := range sitemapNavs {
		if nav.Layout != "" {
			used[nav.Layout] = struct{}{}
		}
	}

	var diags []diagnostic.Diagnostic
	for _, l := range layouts {
		if _, ok := used[l.Name]; !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:    l.File,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelWarning,
				Message: fmt.Sprintf("[TM-13] layout %q is defined but never used by any page, defaultLayout, or sitemap nav block", l.Name),
				Advice:  fmt.Sprintf("Add data-layout=%q to a page or a sitemap <nav data-sitemap> block, set it as defaultLayout in manifest.yaml, or remove layouts/%s.html", l.Name, l.Name),
			})
		}
	}
	return diags
}
