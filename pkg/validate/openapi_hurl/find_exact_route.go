//ff:func feature=validate type=util control=iteration dimension=1 topic=scenario-check
//ff:what findExactRoute — path+method 모두 일치하는 라우트 탐색

package openapi_hurl

import "strings"

// findExactRoute returns the first route whose segments and method both match.
// Returns nil if no exact match is found.
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
