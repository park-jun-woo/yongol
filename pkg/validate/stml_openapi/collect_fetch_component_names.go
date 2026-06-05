//ff:func feature=validate type=util control=iteration dimension=4 topic=stml-openapi
//ff:what collectFetchComponentNames — FetchBlock에서 컴포넌트 이름 재귀 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectFetchComponentNames gathers component names from a FetchBlock's direct
// ComponentRefs, its descendant data-each blocks, its nested fetches, and its
// children (recursively).
func collectFetchComponentNames(f stml.FetchBlock, out map[string]struct{}) {
	for _, c := range f.Components {
		out[c.Name] = struct{}{}
	}
	for _, e := range f.Eaches {
		collectEachComponentNames(e, out)
	}
	for _, nf := range f.NestedFetches {
		collectFetchComponentNames(nf, out)
	}
	for _, ch := range f.Children {
		collectChildComponentNames(ch, out)
	}
}
