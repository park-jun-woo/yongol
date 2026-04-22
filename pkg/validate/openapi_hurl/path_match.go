//ff:func feature=validate type=util control=sequence topic=scenario-check
//ff:what matchHurlEntry — Hurl entry 를 OpenAPI routes 에 매칭하는 공통 helper

package openapi_hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// matchHurlEntry normalizes a hurl entry's path and locates the first
// path-level match (any method) and — when the method also matches — the
// exact route. Return contract:
//
//	pathIdx >= 0  : some OpenAPI route shares the same path shape
//	route  != nil : pathIdx >= 0 AND entry.Method matches that route
//
// Callers (XOH-35/36/37) decide the severity based on which fields are set.
func matchHurlEntry(entry hurl.HurlEntry, routes []apiRoute) (pathIdx int, route *apiRoute) {
	segs := normalizeHurlPath(entry.Path)
	pathIdx = findPathMatch(segs, routes)
	if pathIdx < 0 {
		return -1, nil
	}
	route = findExactRoute(segs, entry.Method, routes)
	return pathIdx, route
}
