//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what collectInvalidTransitions — collects @state transitions in SSaC functions that do not match any diagram event

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// collectInvalidTransitions returns diagnostics for @state sequences whose
// transition is not a valid event in the referenced diagram.
func collectInvalidTransitions(fn ssac.ServiceFunc, diagramByID map[string]*statemachine.StateDiagram) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, seq := range fn.Sequences {
		if seq.Type != "state" {
			continue
		}
		d, ok := diagramByID[seq.DiagramID]
		if !ok {
			continue
		}
		if len(d.ValidFromStates(seq.Transition)) > 0 {
			continue
		}
		file := fn.FileName
		if file == "" {
			file = "ssac/" + fn.Name + ".ssac"
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    file,
			Line:    seq.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XMS-25] transition \"" + seq.Transition + "\" is not a valid event in diagram \"" + seq.DiagramID + "\"",
			Advice:  "Define transition '" + seq.Transition + "' in the stateDiagram",
		})
	}
	return diags
}
