//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what navRoutePatterns — 페이지명 → 해석 라우트 패턴(stml.RoutePaths 첫 패턴) 맵 구성 (레이아웃 data-nav 치환용)

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// navRoutePatterns maps each STML page name to its resolved route pattern
// (stml.RoutePaths first pattern — the Phase001 single source of route
// planning). Layout data-nav page-name references substitute into these
// patterns (page-flow Phase010), the same table the page emitter receives
// as GenerateOptions.RoutePatterns for data-link/data-redirect.
func navRoutePatterns(pages []stml.PageSpec) map[string]string {
	out := map[string]string{}
	for _, p := range pages {
		if paths := stml.RoutePaths(p); len(paths) > 0 {
			out[p.Name] = paths[0]
		}
	}
	return out
}
