//ff:func feature=gen-hurl type=util control=iteration dimension=3
//ff:what buildTransitionOrder — stateDiagram BFS로 event → order idx 계산 (self-loop/terminal 포함)

package hurl

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// buildTransitionOrder walks each stateDiagram BFS from InitialState
// and assigns an increasing index to each event in execution-friendly
// order.
//
// BUG-016 / Phase004 — the previous implementation skipped self-loop
// transitions (t.To == t.From) and transitions that led to an already-
// visited state, leaving events like ExecuteWorkflow (active → active)
// outside the order map. BranchSkip then dropped them from smoke. The
// new walk emits every outgoing transition of each discovered state in
// this priority:
//
//  1. Self-loops first — these preserve state, so emitting them before
//     branch transitions keeps later steps usable.
//  2. Transitions to non-terminal states next — follow-on steps stay
//     executable.
//  3. Transitions to terminal states last — terminal here means a state
//     with no outgoing transitions of its own (e.g. archived).
//
// The walk still does a true BFS over reachable states so multi-stage
// transitions (draft → active → archived) pick up their activation
// prerequisite before the archival step.
func buildTransitionOrder(diagrams []*statemachine.StateDiagram) map[string]int {
	order := map[string]int{}
	idx := 0
	for _, d := range diagrams {
		if d == nil || d.InitialState == "" {
			continue
		}
		terminals := terminalStates(d)
		visited := map[string]bool{d.InitialState: true}
		queue := []string{d.InitialState}
		for len(queue) > 0 {
			state := queue[0]
			queue = queue[1:]
			outs := outgoingByPriority(d, state, terminals)
			for _, t := range outs {
				if _, exists := order[t.Event]; !exists {
					order[t.Event] = idx
					idx++
				}
				if !visited[t.To] {
					visited[t.To] = true
					queue = append(queue, t.To)
				}
			}
		}
	}
	return order
}
