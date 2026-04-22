//ff:func feature=gen-hurl type=util control=iteration dimension=3
//ff:what buildBranchSkipSet — 같은 from 상태의 분기 전이 중 첫째만 유지, 나머지는 skip
package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// buildBranchSkipSet returns events that should be skipped in smoke.
// For each state with multiple outgoing transitions, keep the one with
// the lowest transitionOrder; mark the rest as skip.
func buildBranchSkipSet(diagrams []*statemachine.StateDiagram, transitionOrder map[string]int) map[string]bool {
	skip := map[string]bool{}
	for _, d := range diagrams {
		if d == nil {
			continue
		}
		fromGroups := map[string][]string{}
		for _, t := range d.Transitions {
			fromGroups[t.From] = append(fromGroups[t.From], t.Event)
		}
		for _, events := range fromGroups {
			if len(events) <= 1 {
				continue
			}
			bestEvent := ""
			bestOrder := int(^uint(0) >> 1)
			for _, e := range events {
				if ord, ok := transitionOrder[e]; ok && ord < bestOrder {
					bestOrder = ord
					bestEvent = e
				}
			}
			for _, e := range events {
				if e != bestEvent {
					skip[e] = true
				}
			}
		}
	}
	return skip
}
