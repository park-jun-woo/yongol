//ff:func feature=validate type=util control=sequence topic=stml-openapi
//ff:what collectChildComponentNames — ChildNode에서 컴포넌트 이름 재귀 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectChildComponentNames gathers component names from a ChildNode: a direct
// ComponentRef, or by recursing into a nested fetch, each, action, or state.
func collectChildComponentNames(ch stml.ChildNode, out map[string]struct{}) {
	if ch.Component != nil {
		out[ch.Component.Name] = struct{}{}
	}
	if ch.Fetch != nil {
		collectFetchComponentNames(*ch.Fetch, out)
	}
	if ch.Each != nil {
		collectEachComponentNames(*ch.Each, out)
	}
	if ch.Action != nil {
		collectActionComponentNames(*ch.Action, out)
	}
	if ch.State != nil {
		collectStateComponentNames(*ch.State, out)
	}
}
