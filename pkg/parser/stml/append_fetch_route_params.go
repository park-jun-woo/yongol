//ff:func feature=stml-parse type=util control=iteration dimension=1
//ff:what appendFetchRouteParams — FetchBlock과 중첩 fetch의 route 파라미터를 필수로 누적
package stml

// appendFetchRouteParams appends the route params of a fetch block and its
// nested fetches as required params: a page cannot render without the
// params its data-fetch blocks consume.
func appendFetchRouteParams(params []routeParam, seen map[string]bool, f FetchBlock) []routeParam {
	params = appendRouteParams(params, seen, f.Params, true)
	for _, nested := range f.NestedFetches {
		params = appendFetchRouteParams(params, seen, nested)
	}
	return params
}
