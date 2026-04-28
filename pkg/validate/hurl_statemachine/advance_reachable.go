//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-statemachine
//ff:what advanceReachable — op 에 해당하는 전이의 To 상태를 reachable 에 추가

package hurl_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

// advanceReachable marks every To state that op can legally transition
// into as reachable. Keeps inspectDiagramEntry at depth 2 by extracting
// the secondary loop.
func advanceReachable(d *statemachine.StateDiagram, op string, reachable map[string]bool) {
	for _, t := range d.Transitions {
		if t.Event == op {
			reachable[t.To] = true
		}
	}
}
