//ff:func feature=validate type=helper control=iteration dimension=2 topic=stml-statemachine
//ff:what findRedirectTargetPage — data-redirect 경로가 해석되는 첫 STML 페이지 반환

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// findRedirectTargetPage returns the first STML page whose route patterns
// (stml.RoutePaths) match the given static redirect path, or nil when the
// path resolves to no page (TM-26 in stml_openapi reports that case).
func findRedirectTargetPage(path string, pages []stml.PageSpec) *stml.PageSpec {
	for i := range pages {
		for _, pattern := range stml.RoutePaths(pages[i]) {
			if stml.RouteMatchesPath(pattern, path) {
				return &pages[i]
			}
		}
	}
	return nil
}
