//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what collectRouteParams — 페이지가 소비하는 route.<Name>을 등장 순서·필수(fetch)/선택(action) 구분으로 수집
package stml

// collectRouteParams gathers the route.<Name> params a page consumes, in
// first-appearance order: fetch blocks first (required), then page-level
// actions, then child actions (optional). The traversal order mirrors the
// react emitter's collectAllParams (pkg/generate/react/stml) so the
// derived route segments line up with the emitted useParams()
// destructuring names.
func collectRouteParams(p PageSpec) []routeParam {
	var params []routeParam
	seen := map[string]bool{}
	for _, f := range p.Fetches {
		params = appendFetchRouteParams(params, seen, f)
	}
	for _, a := range p.Actions {
		params = appendRouteParams(params, seen, a.Params, false)
	}
	for _, a := range CollectChildActions(p.Children) {
		params = appendRouteParams(params, seen, a.Params, false)
	}
	return params
}
