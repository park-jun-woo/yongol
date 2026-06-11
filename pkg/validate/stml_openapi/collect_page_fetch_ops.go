//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectPageFetchOps — 페이지의 모든 data-fetch operationId(중첩 포함) 집합 추출

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectPageFetchOps returns the set of every data-fetch operationId declared
// on a page, including nested data-fetch blocks. This mirrors the react
// emitter's collectFetchOps scope — exactly the fetch data variables in scope
// for a form's data-prefill wiring on the same page.
func collectPageFetchOps(page stml.PageSpec) map[string]bool {
	ops := make(map[string]bool)
	for _, f := range page.Fetches {
		addFetchOps(f, ops)
	}
	return ops
}
