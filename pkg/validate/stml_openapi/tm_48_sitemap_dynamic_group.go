//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-openapi
//ff:what TM-48 — 사이트맵 동적 메뉴 그룹 위반: data-entry 블록 선언(비로그인 fetch 성립 불가) / 필수 어휘(fetch·each·link) 누락 (ERROR)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm48SitemapDynamicGroup validates the structure of every sitemap
// dynamic-group declaration (plans/stml/sitemap Phase007): (a) a dynamic
// group inside a data-entry block is an ERROR — the public entry layout
// renders for signed-out visitors, where the (typically protected) list
// fetch can never be satisfied, so the declaration is dead by
// construction; (b) a group declaring any of the vocabulary must declare
// data-fetch, data-each and data-link — without them no item can be
// fetched, enumerated or linked (a missing data-label-field is TM-30's
// finding, the same rule that validates the field itself). Field-level
// judgments (operationId, array field, link target, params, label field)
// belong to the TM-01/07/30/31/32 sitemap extensions.
func tm48SitemapDynamicGroup(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for i, nav := range fs.Sitemap.Navs {
		var entries []sitemapEntry
		appendSitemapEntries(nav.Items, fmt.Sprintf("nav[%d]", i), &entries)
		for _, e := range entries {
			vocab := sitemapDynamicVocab(e.Node)
			if len(vocab) == 0 {
				continue
			}
			if nav.Entry {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fs.Sitemap.FileName,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[TM-48] dynamic menu group at %s sits in a data-entry block — the public entry layout renders for signed-out visitors, so its list fetch can never be satisfied there", e.Path),
					Advice:  "Move the dynamic group into a signed-in layout's sitemap block; data-entry blocks keep static entries only",
				})
			}
			var missing []string
			if e.Node.Fetch == "" {
				missing = append(missing, "data-fetch")
			}
			if e.Node.Each == "" {
				missing = append(missing, "data-each")
			}
			if e.Node.Link == "" {
				missing = append(missing, "data-link")
			}
			if len(missing) > 0 {
				diags = append(diags, diagnostic.Diagnostic{
					File:    fs.Sitemap.FileName,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: fmt.Sprintf("[TM-48] dynamic menu group at %s declares %s but is missing %s — the group cannot fetch, enumerate or link its items", e.Path, strings.Join(vocab, "/"), strings.Join(missing, ", ")),
					Advice:  "Declare data-fetch (list operationId), data-each (response array field) and data-link (item target page) together on the group's <ul>",
				})
			}
		}
	}
	return diags
}
