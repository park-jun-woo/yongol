//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what addEntryRoots — data-entry 블록의 전 깊이 페이지를 (실존하는 것만) 루트로 편입

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// addEntryRoots marks every page of a data-entry nav block as a
// reachability root — the block declares public entry points users land
// on directly (DESIGN §4.10 root set). Names matching no STML page are
// dropped (TM-39's finding).
func addEntryRoots(items []stml.SitemapNode, names map[string]bool, g *pageGraph) {
	for _, name := range sitemapNavPages(items) {
		if names[name] {
			g.Roots[name] = true
		}
	}
}
