//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what findPathMatch — method 무시하고 segs 만 일치하는 첫 route index

package hurl_openapi

// findPathMatch returns the index of the first route whose segments
// match, ignoring method. -1 means no path match exists.
func findPathMatch(segs []string, routes []apiRoute) int {
	for i := range routes {
		if segmentsMatch(segs, routes[i].Segments) {
			return i
		}
	}
	return -1
}
