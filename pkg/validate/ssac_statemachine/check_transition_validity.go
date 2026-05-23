//ff:func feature=validate type=util control=sequence topic=states
//ff:what checkTransitionValidity — 단일 @state 시퀀스의 전이가 다이어그램 이벤트에 유효한지 검사

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// checkTransitionValidity returns a diagnostic and true if the sequence's
// transition is invalid in the referenced diagram.
func checkTransitionValidity(fn ssac.ServiceFunc, seq ssac.Sequence, diagramByID map[string]*statemachine.StateDiagram) (diagnostic.Diagnostic, bool) {
	if seq.Type != "state" {
		return diagnostic.Diagnostic{}, false
	}
	d, ok := diagramByID[seq.DiagramID]
	if !ok {
		return diagnostic.Diagnostic{}, false
	}
	if len(d.ValidFromStates(seq.Transition)) > 0 {
		return diagnostic.Diagnostic{}, false
	}
	file := fn.FileName
	if file == "" {
		file = "ssac/" + fn.Name + ".ssac"
	}
	return diagnostic.Diagnostic{
		File:        file,
		Line:        seq.Line,
		Phase:       diagnostic.PhaseValidate,
		Level:       diagnostic.LevelError,
		Message:     "[XMS-25] transition \"" + seq.Transition + "\" is not a valid event in diagram \"" + seq.DiagramID + "\"",
		Advice:      "Define transition '" + seq.Transition + "' in the stateDiagram",
		OperationID: fn.Name,
	}, true
}
