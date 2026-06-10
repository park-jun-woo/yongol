//ff:func feature=validate type=helper control=iteration dimension=1 topic=stml-statemachine
//ff:what transitionToStates — 주어진 이벤트로 라벨된 전이의 도착 상태 집합 반환

package stml_statemachine

import "github.com/park-jun-woo/yongol/pkg/parser/statemachine"

// transitionToStates returns the arrival states of every transition in the
// diagram labeled with the given event. An empty result means the event is
// not a transition label of this diagram (TM-23 then stays silent).
func transitionToStates(d *statemachine.StateDiagram, event string) []string {
	var out []string
	for _, t := range d.Transitions {
		if t.Event == event {
			out = append(out, t.To)
		}
	}
	return out
}
