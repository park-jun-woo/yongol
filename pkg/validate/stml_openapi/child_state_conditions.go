//ff:func feature=validate type=util control=selection dimension=1 topic=stml-openapi
//ff:what 단일 ChildNode에서 data-state 조건을 수집하고 하위 블록으로 재귀한다

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// childStateConditions returns the data-state conditions contained in a single
// child node, recursing into the children of fetch, each, and state blocks.
func childStateConditions(c stml.ChildNode) []string {
	switch c.Kind {
	case "state":
		return append([]string{c.State.Condition}, collectStateConditions(c.State.Children)...)
	case "fetch":
		return collectStateConditions(c.Fetch.Children)
	case "each":
		return collectStateConditions(c.Each.Children)
	default:
		return nil
	}
}
