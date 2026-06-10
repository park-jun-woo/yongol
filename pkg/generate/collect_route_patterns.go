//ff:func feature=generate type=util control=iteration dimension=1
//ff:what collectRoutePatterns — 페이지명 → 해석 라우트 패턴(stml.RoutePaths 첫 패턴) 맵 구성

package generate

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectRoutePatterns maps each STML page name to its resolved route
// pattern (stml.RoutePaths first pattern) — the same table the App.tsx
// router uses (Phase001 single source). data-link emission substitutes
// link param sources into these patterns (page-flow Phase007).
func collectRoutePatterns(pages []stmlparser.PageSpec) map[string]string {
	routePatterns := map[string]string{}
	for _, p := range pages {
		if paths := stmlparser.RoutePaths(p); len(paths) > 0 {
			routePatterns[p.Name] = paths[0]
		}
	}
	return routePatterns
}
