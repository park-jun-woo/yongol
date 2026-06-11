//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-openapi
//ff:what TM-42 — data-index 정합 위반: 2개 이상 / page 없는 노드 / 필수 파라미터 라우트(TM-34 판정) / manifest.frontend.index 모순 (ERROR)

package stml_openapi

import (
	"fmt"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// tm42SitemapIndexConflict validates the sitemap's data-index declaration
// (plans/stml/sitemap Phase001): at most one data-index across the whole
// sitemap — "/" redirects to exactly one place — plus the per-entry checks
// in tm42IndexEntryDiags (entry without data-page, required-parameter
// route, manifest.frontend.index contradiction).
func tm42SitemapIndexConflict(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var indexEntries []sitemapEntry
	for _, e := range collectSitemapEntries(fs.Sitemap) {
		if e.Node.Index {
			indexEntries = append(indexEntries, e)
		}
	}
	if len(indexEntries) == 0 {
		return nil
	}

	var diags []diagnostic.Diagnostic
	if len(indexEntries) > 1 {
		var paths []string
		for _, e := range indexEntries {
			paths = append(paths, e.Path)
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    fs.Sitemap.FileName,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: fmt.Sprintf("[TM-42] data-index is declared %d times in the sitemap (at %s) — \"/\" can redirect to exactly one page", len(indexEntries), strings.Join(paths, ", ")),
			Advice:  "Keep data-index on exactly one <li data-page> entry",
		})
	}
	for _, e := range indexEntries {
		diags = append(diags, tm42IndexEntryDiags(fs, e)...)
	}
	return diags
}
