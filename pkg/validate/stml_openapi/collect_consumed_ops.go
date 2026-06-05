//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectConsumedOps — STML data-fetch/data-action + 컴포넌트 .tsx의 api 호출에서 소비된 operationId 집합 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectConsumedOps returns the set of operationIds consumed across all STML
// pages: data-fetch and data-action blocks, plus api.<operationId>( calls
// inside referenced components' .tsx files under
// <specsDir>/frontend/components/. Component candidates are filtered against
// ops (the set of real OpenAPI operationIds) to avoid false positives. A zero
// specsDir or empty ops skips the component scan, leaving behavior identical to
// the fetch/action-only collection.
func collectConsumedOps(pages []stml.PageSpec, specsDir string, ops map[string]struct{}) map[string]struct{} {
	out := make(map[string]struct{})
	for _, page := range pages {
		for _, f := range page.Fetches {
			collectFetchOps(f, out)
		}
		for _, a := range page.Actions {
			out[a.OperationID] = struct{}{}
		}
	}
	names := collectComponentNames(pages)
	collectComponentApiOps(names, specsDir, ops, out)
	return out
}
