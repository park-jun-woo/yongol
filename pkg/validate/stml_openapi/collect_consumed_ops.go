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

func collectChildOps(ch stml.ChildNode, out map[string]struct{}) {
	if ch.Action != nil {
		out[ch.Action.OperationID] = struct{}{}
	}
	if ch.Fetch != nil {
		collectFetchOps(*ch.Fetch, out)
	}
	if ch.State != nil {
		for _, sc := range ch.State.Children {
			collectChildOps(sc, out)
		}
	}
}
