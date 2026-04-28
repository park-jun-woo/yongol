//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what findExactRoute — segs/method 가 모두 일치하는 첫 route 반환

package hurl_openapi

import "strings"

// findExactRoute returns the first route whose segments and method both
// match the given hurl entry. Returns nil when no operation matches.
func findExactRoute(segs []string, method string, routes []apiRoute) *apiRoute {
	m := strings.ToUpper(method)
	for i := range routes {
		if routes[i].Method != m {
			continue
		}
		if segmentsMatch(segs, routes[i].Segments) {
			return &routes[i]
		}
	}
	return nil
}
