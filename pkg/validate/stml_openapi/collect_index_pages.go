//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectIndexPages — "/" 가 도달하는 인덱스 페이지명 수집 (manifest.frontend.index ∪ sitemap data-index ∪ data-route="/" 마운트)

package stml_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// collectIndexPages names the existing pages "/" reaches, deduplicated:
// manifest.frontend.index, sitemap data-index entries, and pages mounting
// "/" directly via data-route — the three index vehicles TM-34/35/42
// already arbitrate. buildPageGraph seeds the reachability roots with
// them and resolveRedirectTargets resolves data-redirect="/" to them.
// Names that match no STML page are dropped (TM-34/39's findings).
func collectIndexPages(fs *yongol.Fullstack) []string {
	var candidates []string
	if fs.Manifest != nil {
		candidates = append(candidates, fs.Manifest.Frontend.Index)
	}
	if fs.Sitemap != nil {
		for _, e := range collectSitemapEntries(fs.Sitemap) {
			if e.Node.Index {
				candidates = append(candidates, e.Node.Page)
			}
		}
	}
	for _, p := range fs.STMLPages {
		for _, pattern := range stml.RoutePaths(p) {
			if pattern == "/" {
				candidates = append(candidates, p.Name)
				break
			}
		}
	}

	seen := make(map[string]bool, len(candidates))
	var out []string
	for _, name := range candidates {
		if name == "" || seen[name] || findPageByName(fs.STMLPages, name) == nil {
			continue
		}
		seen[name] = true
		out = append(out, name)
	}
	return out
}
