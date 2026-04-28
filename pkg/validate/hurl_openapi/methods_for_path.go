//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what methodsForPath — segs 와 일치하는 route 의 HTTP method 정렬 리스트

package hurl_openapi

import "sort"

// methodsForPath returns the sorted list of HTTP methods declared for
// the first route whose segments match.
func methodsForPath(segs []string, routes []apiRoute) []string {
	var out []string
	for _, r := range routes {
		if segmentsMatch(segs, r.Segments) {
			out = append(out, r.Method)
		}
	}
	sort.Strings(out)
	return out
}
