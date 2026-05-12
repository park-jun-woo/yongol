//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectConsumedOps — STML 전 페이지의 data-fetch/data-action에서 참조하는 operationId 집합 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectConsumedOps returns the set of operationIds referenced by
// all data-fetch and data-action blocks across all STML pages.
func collectConsumedOps(pages []stml.PageSpec) map[string]struct{} {
	out := make(map[string]struct{})
	for _, page := range pages {
		for _, f := range page.Fetches {
			collectFetchOps(f, out)
		}
		for _, a := range page.Actions {
			out[a.OperationID] = struct{}{}
		}
	}
	return out
}
