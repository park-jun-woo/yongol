//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what appendEdgeOnce — 페이지 그래프에 간선 추가 (동일 간선 중복 억제)

package stml_openapi

// appendEdgeOnce adds one edge to the page graph unless it already
// exists. The breadcrumb up-edges (DESIGN §4.10 edge (d), Phase004) are
// collected after the data-link/data-redirect loop and frequently
// duplicate it — a detail page typically both redirects to its list and
// links up to it through the breadcrumb — and a duplicate edge would only
// distort the BFS frontier and the Edges fixtures, never the result.
func appendEdgeOnce(g *pageGraph, from, to string) {
	for _, t := range g.Edges[from] {
		if t == to {
			return
		}
	}
	g.Edges[from] = append(g.Edges[from], to)
}
