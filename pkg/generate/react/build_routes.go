//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what STML PageSpec 목록을 정렬된 stmlRoute 정의로 변환한다 (레이아웃 배정 3단 사슬: 페이지 > sitemap 블록 > defaultLayout)

package react

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// buildRoutes converts STML PageSpecs into sorted route definitions.
// Layout assignment is the three-step chain of plans/stml/sitemap
// Phase003 (DESIGN §4.9 — the specific beats the general): the page's own
// data-layout, then the data-layout of the sitemap nav block listing the
// page (sitemapLayouts — sitemapPageLayouts, keyed by page name; nil
// without a sitemap, restoring the original two-step chain), then
// defaultLayout. protectedPages flags pages (by FileName) that consume a
// security-protected OpenAPI op (Phase005 — resolveProtectedPages); their
// routes get Protected=true so the renderer wraps them with
// <ProtectedRoute>. A nil map leaves every route public.
func buildRoutes(pages []stml.PageSpec, defaultLayout string, protectedPages map[string]bool, sitemapLayouts map[string]string) []stmlRoute {
	routes := make([]stmlRoute, 0, len(pages))
	for _, p := range pages {
		rs := pageToRoutes(p)
		resolvedLayout := p.Layout
		if resolvedLayout == "" {
			resolvedLayout = sitemapLayouts[p.Name]
		}
		if resolvedLayout == "" {
			resolvedLayout = defaultLayout
		}
		for i := range rs {
			rs[i].Layout = resolvedLayout
			rs[i].Protected = protectedPages[p.FileName]
		}
		routes = append(routes, rs...)
	}
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Path < routes[j].Path
	})
	return routes
}
