//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectFetchOps — FetchBlock 에서 재귀적으로 operationId 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectFetchOps recursively collects operationIds from a FetchBlock,
// its nested fetches, and actions inside data-state children.
func collectFetchOps(f stml.FetchBlock, out map[string]struct{}) {
	out[f.OperationID] = struct{}{}
	for _, child := range f.NestedFetches {
		collectFetchOps(child, out)
	}
	for _, child := range f.Children {
		collectChildOps(child, out)
	}
}
