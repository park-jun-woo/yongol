//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what findBridgeEvent — currentState에서 target 집합으로 이어지는 1-hop transition event 탐색

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// findBridgeEvent searches for a single transition that starts at
// currentState and ends in one of the targetStates. Returns the event
// name and the target state on success, or empty strings when no
// direct bridge exists. The search is one-hop only — multi-hop path
// finding would complicate the smoke walker without clear benefit:
// real projects keep the state graph shallow by design.
func findBridgeEvent(diagrams []*statemachine.StateDiagram, currentState string, targetStates map[string]bool) (string, string) {
	for _, d := range diagrams {
		if d == nil {
			continue
		}
		for _, t := range d.Transitions {
			if t.From != currentState {
				continue
			}
			if targetStates[t.To] {
				return t.Event, t.To
			}
		}
	}
	return "", ""
}
