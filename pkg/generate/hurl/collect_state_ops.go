//ff:func feature=gen-hurl type=util control=iteration dimension=2
//ff:what collectStateOps — stateDiagram event 로 사용된 operationID 집합 수집
package hurl

import "github.com/park-jun-woo/yongol/pkg/yongol"

// collectStateOps returns the set of operationIds that appear as transition
// events in any StateDiagram. Used by buildCreateSteps to skip state-transition
// POSTs so that buildStateTransitions is the sole emitter (with BFS order and
// branch-skip applied).
func collectStateOps(fs *yongol.Fullstack) map[string]bool {
	set := map[string]bool{}
	if fs == nil {
		return set
	}
	for _, d := range fs.StateDiagrams {
		if d == nil {
			continue
		}
		for _, tr := range d.Transitions {
			if tr.Event != "" {
				set[tr.Event] = true
			}
		}
	}
	return set
}
