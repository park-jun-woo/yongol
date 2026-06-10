//ff:func feature=validate type=helper control=iteration dimension=1 topic=stml-statemachine
//ff:what collectPageStateConditions — 페이지 수준 data-state 가드 조건 수집 (each·중첩 state 내부 제외)

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/stml"

// collectPageStateConditions returns the page-level data-state guard
// conditions of a page: state binds reachable through fetch and static
// wrappers. It does not descend into each items (per-item guards) or into
// the conditional content of a state bind itself — those are sub-guards,
// not the guards that gate what the page shows on arrival (TM-23).
func collectPageStateConditions(children []stml.ChildNode) []string {
	var out []string
	for _, c := range children {
		switch c.Kind {
		case "state":
			out = append(out, c.State.Condition)
		case "fetch":
			out = append(out, collectPageStateConditions(c.Fetch.Children)...)
		case "static":
			out = append(out, collectPageStateConditions(c.Static.Children)...)
		}
	}
	return out
}
