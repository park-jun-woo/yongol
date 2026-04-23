//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what eventFromStates — event 를 유발할 수 있는 from-state 집합 수집

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// eventFromStates returns the set of states from which `event` is a
// declared transition. A single event may originate from several
// states (e.g. ActivateWorkflow fires from both `draft` and `paused`).
func eventFromStates(diagrams []*statemachine.StateDiagram, event string) map[string]bool {
	fromStates := map[string]bool{}
	for _, d := range diagrams {
		if d == nil {
			continue
		}
		for _, t := range d.Transitions {
			if t.Event == event {
				fromStates[t.From] = true
			}
		}
	}
	return fromStates
}
