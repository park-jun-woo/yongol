//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what reachablePages — 페이지 그래프 루트에서 BFS 로 도달 가능한 페이지 집합 계산

package stml_openapi

// reachablePages runs the BFS of DESIGN §4.10 over the page graph: the
// frontier starts at the roots (index/entry pages plus menu-rendered
// sitemap entries) and follows data-link/data-redirect edges. Seeding
// iterates g.Pages so ghost root names (sitemap entries naming no page —
// TM-39's finding) never enter the set and the order stays deterministic.
func reachablePages(g *pageGraph) map[string]bool {
	reached := make(map[string]bool, len(g.Pages))
	var queue []string
	for _, name := range g.Pages {
		if g.Roots[name] {
			reached[name] = true
			queue = append(queue, name)
		}
	}
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]
		for _, next := range g.Edges[cur] {
			if !reached[next] {
				reached[next] = true
				queue = append(queue, next)
			}
		}
	}
	return reached
}
