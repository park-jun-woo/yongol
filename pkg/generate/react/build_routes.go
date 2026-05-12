//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what STML PageSpec 목록을 정렬된 stmlRoute 정의로 변환한다

package react

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// buildRoutes converts STML PageSpecs into sorted route definitions.
// defaultLayout is applied to pages that have no explicit Layout set.
func buildRoutes(pages []stml.PageSpec, defaultLayout string) []stmlRoute {
	routes := make([]stmlRoute, 0, len(pages))
	for _, p := range pages {
		rs := pageToRoutes(p)
		resolvedLayout := p.Layout
		if resolvedLayout == "" {
			resolvedLayout = defaultLayout
		}
		for i := range rs {
			rs[i].Layout = resolvedLayout
		}
		routes = append(routes, rs...)
	}
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})
	return routes
}
