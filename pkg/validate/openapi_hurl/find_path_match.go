//ff:func feature=validate type=util control=iteration dimension=1 topic=scenario-check
//ff:what findPathMatch — 세그먼트가 일치하는 첫 라우트 인덱스 (없으면 -1)

package openapi_hurl

// findPathMatch returns the index of the first route whose segments match,
// or -1 if none match.
func findPathMatch(segs []string, routes []apiRoute) int {
	for i := range routes {
		if segmentsMatch(segs, routes[i].Segments) {
			return i
		}
	}
	return -1
}
