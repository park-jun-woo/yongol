//ff:func feature=gen-react type=util control=iteration dimension=1
//ff:what 페이지의 fetch/action 파라미터에 route. 접두사가 있는지 확인한다

package react

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// pageHasRouteParam returns true if any fetch or action block in the page
// contains a data-param-* attribute with a "route." prefixed source.
func pageHasRouteParam(p stml.PageSpec) bool {
	for _, f := range p.Fetches {
		if hasRouteSource(f.Params) {
			return true
		}
	}
	for _, a := range p.Actions {
		if hasRouteSource(a.Params) {
			return true
		}
	}
	return false
}
