//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-40 — 같은 페이지가 사이트맵 전체에 2회 이상 등장 (canonical 위치 위반, 두 위치 표기, ERROR)

package stml_openapi

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm40SitemapDuplicatePage enforces the canonical-position rule
// (plans/stml/sitemap Phase001, DESIGN §4.3): a page appears at most once
// across the whole sitemap, nav blocks included — breadcrumbs and active
// highlighting require a unique parent. The message names both positions so
// the fix needs no search. Cross-references belong to data-link.
func tm40SitemapDuplicatePage(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	firstAt := make(map[string]string)
	var diags []diagnostic.Diagnostic
	for _, e := range collectSitemapEntries(fs.Sitemap) {
		if e.Node.Page == "" {
			continue
		}
		first, seen := firstAt[e.Node.Page]
		if !seen {
			firstAt[e.Node.Page] = e.Path
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fs.Sitemap.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-40] page %q appears more than once in the sitemap — first at %s, again at %s (a page has exactly one canonical position)", e.Node.Page, first, e.Path),
			Advice:  "Keep one canonical entry; a cross-reference from another screen is data-link's job on the referring page",
		})
	}
	return diags
}
