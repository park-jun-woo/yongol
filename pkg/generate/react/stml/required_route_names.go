//ff:func feature=stml-gen type=util control=iteration dimension=1
//ff:what 페이지의 data-fetch가 소비하는 route.<Name> 세그먼트 집합(필수 라우트 파라미터)을 만든다 (BUG-136)
package stml

import stmlparser "github.com/park-jun-woo/yongol/pkg/parser/stml"

// requiredRouteNames returns the set of route segment names consumed by some
// data-fetch on the page (including nested fetches). These are the required
// (":Name") segments; any route param not in this set is optional. Mirrors
// collectRouteParams's fetch-first required rule.
func requiredRouteNames(page stmlparser.PageSpec) map[string]bool {
	var fetchBinds []stmlparser.ParamBind
	for _, f := range page.Fetches {
		fetchBinds = collectFetchParamBinds(f, fetchBinds)
	}
	set := map[string]bool{}
	for _, b := range fetchBinds {
		if name, ok := routeSegmentName(b.Source); ok {
			set[name] = true
		}
	}
	return set
}
