//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what buildTransitionMap — StateDiagram.Transitions → map[from][event]to (초기 전이 제외)

package state

import "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

// buildTransitionMap converts a StateDiagram's Transitions slice into a nested
// map[currentState]map[event]nextState. Initial transitions from "[*]" are
// excluded because they represent row creation, not runtime transitions.
func buildTransitionMap(d *statemachine.StateDiagram) map[string]map[string]string {
	m := make(map[string]map[string]string)
	for _, t := range d.Transitions {
		if t.From == "[*]" {
			continue
		}
		if m[t.From] == nil {
			m[t.From] = make(map[string]string)
		}
		m[t.From][t.Event] = t.To
	}
	return m
}
