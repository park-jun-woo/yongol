//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-41 — sitemap <nav data-layout> 값이 layouts/ 에 없음 (TM-11 과 동일 판정, ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm41SitemapLayoutNotFound checks every <nav data-sitemap data-layout>
// value against the layouts defined in layouts/ — the same existence
// judgment TM-11 applies to a page's data-layout (plans/stml/sitemap
// Phase001). Navs without data-layout delegate to defaultLayout and are
// skipped.
func tm41SitemapLayoutNotFound(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	names := make(map[string]struct{}, len(fs.Layouts))
	for _, l := range fs.Layouts {
		names[l.Name] = struct{}{}
	}

	var diags []diagnostic.Diagnostic
	for i, nav := range fs.Sitemap.Navs {
		if nav.Layout == "" {
			continue
		}
		if _, ok := names[nav.Layout]; !ok {
			diags = append(diags, diagnostic.Diagnostic{
				File:    fs.Sitemap.FileName,
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: fmt.Sprintf("[TM-41] sitemap data-layout %q on nav[%d] not found in layouts/", nav.Layout, i),
				Advice:  fmt.Sprintf("Create layouts/%s.html or fix the data-layout attribute value", nav.Layout),
			})
		}
	}
	return diags
}
