//ff:func feature=validate type=rule control=sequence topic=hurl-statemachine
//ff:what inspectDiagramEntry — 한 entry 가 diagram 전이에 해당하는지 보고 reachable 집합 갱신

package hurl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// inspectDiagramEntry examines a single hurl entry in the context of a
// diagram. When the entry's operationId matches one of the diagram's
// transitions and that transition's From state is not yet reachable,
// emits an XOH-05 WARNING. Otherwise updates reachable with any To
// states reached by matching transitions.
func inspectDiagramEntry(e hurl.HurlEntry, opID map[string]string, d *statemachine.StateDiagram, reachable map[string]bool) []diagnostic.Diagnostic {
	key := e.Method + " " + normPath(e.Path, nil, nil)
	op := opID[key]
	if op == "" {
		return nil
	}
	from := transitionFrom(d, op)
	if from == "" {
		return nil
	}
	if !reachable[from] {
		return []diagnostic.Diagnostic{{
			File:    e.File,
			Line:    e.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[XOH-05] call " + op + " requires state " + from + " which was not reached by prior hurl steps",
			Advice:  "Add the preceding transition(s) declared in states/" + d.ID + ".md before " + op + ", or accept this as a negative-path scenario",
		}}
	}
	advanceReachable(d, op, reachable)
	return nil
}
