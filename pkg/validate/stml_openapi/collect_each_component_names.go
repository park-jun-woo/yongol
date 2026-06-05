//ff:func feature=validate type=util control=iteration dimension=2 topic=stml-openapi
//ff:what collectEachComponentNames — EachBlock에서 컴포넌트 이름 재귀 수집

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectEachComponentNames gathers component names from an EachBlock's direct
// ComponentRefs and its children (recursively).
func collectEachComponentNames(e stml.EachBlock, out map[string]struct{}) {
	for _, c := range e.Components {
		out[c.Name] = struct{}{}
	}
	for _, ch := range e.Children {
		collectChildComponentNames(ch, out)
	}
}
