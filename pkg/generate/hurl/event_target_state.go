//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what eventTargetState — 현재 상태에서 event 실행 시 도달 상태 (해당 없으면 fromState 유지)

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// eventTargetState resolves the state reached after invoking `event`
// when the walker is currently at `fromState`. It prefers a direct
// match where a transition exists from `fromState`; if none exists it
// falls back to the first transition carrying the event name. Returns
// fromState unchanged when the event is unknown.
func eventTargetState(diagrams []*statemachine.StateDiagram, fromState, event string) string {
	fallback := ""
	for _, d := range diagrams {
		if d == nil {
			continue
		}
		for _, t := range d.Transitions {
			if t.Event != event {
				continue
			}
			if t.From == fromState {
				return t.To
			}
			if fallback == "" {
				fallback = t.To
			}
		}
	}
	if fallback != "" {
		return fallback
	}
	return fromState
}
