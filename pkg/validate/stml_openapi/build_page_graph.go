//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what buildPageGraph — 도달성 그래프 구축 (노드=전체 페이지, 루트=인덱스∪entry∪메뉴 렌더, 간선=link/redirect/브레드크럼 상행)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/yongol"

// buildPageGraph assembles the reachability graph of DESIGN §4.10
// (plans/stml/sitemap Phase002). Nodes: every STML page, sitemap listing
// notwithstanding — listing is a node, not an edge. Roots: the index
// pages (collectIndexPages), every page of a data-entry block, and the
// menu-rendered sitemap entries (collectSitemapRoots). Edges: data-link
// and resolvable data-redirect targets (edges (b)/(c)) plus the breadcrumb
// up-edges of Phase004 (edge (d) — child page → MenuRenderable sitemap
// ancestor, collected last so appendEdgeOnce dedupes against the
// link/redirect ones), kept only between existing pages. The sitemap walk
// also records the listing state and per-page menu-block reasons that
// TM-43 turns into its cause classification.
func buildPageGraph(fs *yongol.Fullstack) *pageGraph {
	g := &pageGraph{
		Roots:       map[string]bool{},
		Edges:       map[string][]string{},
		InSitemap:   map[string]bool{},
		MenuBlocked: map[string]string{},
	}
	names := make(map[string]bool, len(fs.STMLPages))
	for _, p := range fs.STMLPages {
		g.Pages = append(g.Pages, p.Name)
		names[p.Name] = true
	}

	indexPages := collectIndexPages(fs)
	for _, name := range indexPages {
		g.Roots[name] = true
	}
	if fs.Sitemap != nil {
		collectSitemapRoots(fs.Sitemap, names, fs.STMLPages, g)
	}

	for _, p := range fs.STMLPages {
		for _, target := range collectPageEdges(p, fs.STMLPages, indexPages) {
			if names[target] {
				g.Edges[p.Name] = append(g.Edges[p.Name], target)
			}
		}
	}
	if fs.Sitemap != nil {
		collectBreadcrumbEdges(fs.Sitemap, names, fs.STMLPages, g)
	}
	return g
}
