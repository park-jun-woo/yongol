//ff:func feature=validate type=util control=iteration dimension=1 topic=hurl-openapi
//ff:what Hurl/OpenAPI 경로 정규화 + 세그먼트 매칭 helpers

package hurl_openapi

import (
	"regexp"
	"strings"
)

var (
	reHurlVar     = regexp.MustCompile(`\{\{.+?\}\}`)
	reHurlNumeric = regexp.MustCompile(`^\d+$`)
	reOpenAPIVar  = regexp.MustCompile(`^\{.+\}$`)
)

// normalizeHurlPath converts a Hurl request path into segments suitable
// for comparison with normalizeOpenAPIPath output. `{{var}}` and pure
// numeric literals collapse to `:param` so path-parameter positions
// match whatever OpenAPI declared.
func normalizeHurlPath(path string) []string {
	path = strings.TrimSpace(path)
	if idx := strings.Index(path, "?"); idx >= 0 {
		path = path[:idx]
	}
	parts := strings.Split(strings.Trim(path, "/"), "/")
	var segs []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		if reHurlVar.MatchString(p) || reHurlNumeric.MatchString(p) {
			segs = append(segs, ":param")
		} else {
			segs = append(segs, p)
		}
	}
	return segs
}

// normalizeOpenAPIPath converts a path declaration like
// `/users/{id}/posts/{postId}` into `["users", ":param", "posts",
// ":param"]`.
func normalizeOpenAPIPath(path string) []string {
	parts := strings.Split(strings.Trim(path, "/"), "/")
	var segs []string
	for _, p := range parts {
		if p == "" {
			continue
		}
		if reOpenAPIVar.MatchString(p) {
			segs = append(segs, ":param")
		} else {
			segs = append(segs, p)
		}
	}
	return segs
}

// segmentsMatch compares two normalized segment slices element-wise.
func segmentsMatch(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

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
