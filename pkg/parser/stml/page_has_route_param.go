//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what 페이지의 fetch/action 파라미터에 route. 접두사 소스가 있는지 확인
package stml

// pageHasRouteParam returns true if any fetch or action block in the page
// contains a data-param-* attribute with a "route." prefixed source.
// Mirrors pkg/generate/react/page_has_route_param.go — kept in sync manually.
func pageHasRouteParam(p PageSpec) bool {
	for _, f := range p.Fetches {
		if paramsHaveRouteSource(f.Params) {
			return true
		}
	}
	for _, a := range p.Actions {
		if paramsHaveRouteSource(a.Params) {
			return true
		}
	}
	return false
}
