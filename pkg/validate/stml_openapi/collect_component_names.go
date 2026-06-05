//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what collectComponentNames — 전 페이지 트리에서 참조된 컴포넌트 이름 집합 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectComponentNames returns the set of component names referenced anywhere
// in the STML page tree. Names appear on four paths: FetchBlock.Components,
// EachBlock.Components, ChildNode.Component, and ActionBlock.Fields entries
// whose Tag has the "data-component:" prefix.
func collectComponentNames(pages []stml.PageSpec) map[string]struct{} {
	out := make(map[string]struct{})
	for _, page := range pages {
		collectPageComponentNames(page, out)
	}
	return out
}
