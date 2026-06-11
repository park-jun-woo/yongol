//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what pageHasGetByIdFetch — 페이지에 route 파라미터를 소비하는 GET fetch(단건 조회)가 있는지 판정

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// pageHasGetByIdFetch reports whether the page declares a GET-by-id fetch: a
// data-fetch served under GET that consumes a route path parameter (a
// data-param-* sourced from "route."). That is the read-current-value half of
// the canonical edit page, the signal TM-55 keys its no-prefill warning on.
func pageHasGetByIdFetch(page stml.PageSpec, opMap map[string]operationEntry) bool {
	for _, f := range page.Fetches {
		entry, ok := opMap[f.OperationID]
		if !ok || entry.method != "GET" {
			continue
		}
		if fetchHasRouteParam(f) {
			return true
		}
	}
	return false
}
