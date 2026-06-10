//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what collectChildOps — ChildNode 에서 action/fetch/state/each/static 재귀적으로 operationId 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

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
	// Row actions inside data-each (page-flow Phase006) and actions nested
	// in static wrappers count as consumption too.
	if ch.Each != nil {
		for _, ec := range ch.Each.Children {
			collectChildOps(ec, out)
		}
	}
	if ch.Static != nil {
		for _, sc := range ch.Static.Children {
			collectChildOps(sc, out)
		}
	}
}
