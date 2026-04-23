//ff:func feature=gen-hurl type=util control=iteration dimension=1
//ff:what terminalStates — stateDiagram에서 나가는 transition 없는 상태 집합 수집

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// terminalStates returns the set of states in d that have no outgoing
// transitions. Used by buildTransitionOrder to push transitions into
// terminal states (archived, closed, ...) to the tail of the emission
// order so smoke can run the full lifecycle before locking the record.
func terminalStates(d *statemachine.StateDiagram) map[string]bool {
	hasOut := map[string]bool{}
	states := map[string]bool{}
	for _, t := range d.Transitions {
		hasOut[t.From] = true
		states[t.From] = true
		states[t.To] = true
	}
	term := map[string]bool{}
	for s := range states {
		if !hasOut[s] {
			term[s] = true
		}
	}
	return term
}
