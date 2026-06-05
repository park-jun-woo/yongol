//ff:func feature=validate type=util control=iteration dimension=1 topic=stml-openapi
//ff:what STML 자식 트리를 순회하며 모든 data-state 조건 문자열을 수집한다 (TM-17용)

package stml_openapi

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectStateConditions walks a list of STML child nodes in DOM order and
// returns every data-state condition string, recursing into fetch, each, and
// state blocks so nested guards are reached.
func collectStateConditions(children []stml.ChildNode) []string {
	var conds []string
	for _, c := range children {
		conds = append(conds, childStateConditions(c)...)
	}
	return conds
}
