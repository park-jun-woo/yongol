//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what PageConsumedOps — 단일 STML 페이지가 소비하는 operationId 집합 (XMO-10 소비 인덱스의 페이지 단위 뷰)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// PageConsumedOps returns the set of operationIds a single STML page
// consumes: its data-fetch blocks (recursively), its data-action blocks,
// plus api.<operationId>( calls inside the components it references under
// <specsDir>/frontend/components/. It is the per-page view of the same
// consumption definition collectConsumedOps unions for XMO-10, exported
// (Phase005) so pkg/generate/react infers protected routes from the exact
// index the validator already trusts. Component candidates are filtered
// against ops (the set of real OpenAPI operationIds); a zero specsDir or
// empty ops skips the component scan.
func PageConsumedOps(page stml.PageSpec, specsDir string, ops map[string]struct{}) map[string]struct{} {
	out := make(map[string]struct{})
	for _, f := range page.Fetches {
		collectFetchOps(f, out)
	}
	for _, a := range page.Actions {
		out[a.OperationID] = struct{}{}
	}
	names := make(map[string]struct{})
	collectPageComponentNames(page, names)
	collectComponentApiOps(names, specsDir, ops, out)
	return out
}
