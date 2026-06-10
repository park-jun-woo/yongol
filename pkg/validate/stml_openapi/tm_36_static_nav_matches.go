//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what tm36StaticNavMatches — 정적 data-nav 경로가 어떤 페이지의 해석 라우트 패턴에든 매칭되는지 판정

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// tm36StaticNavMatches reports whether a "/"-prefixed data-nav path
// matches at least one page's resolved route pattern
// (stml.RouteMatchesPath over stml.RoutePaths — the same matching TM-26
// applies to static data-redirect paths).
func tm36StaticNavMatches(path string, pages []stml.PageSpec) bool {
	for _, p := range pages {
		for _, pattern := range stml.RoutePaths(p) {
			if stml.RouteMatchesPath(pattern, path) {
				return true
			}
		}
	}
	return false
}
