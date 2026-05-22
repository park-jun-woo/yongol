//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what collectMissingDiagrams — collects @state sequences in SSaC functions that reference a non-existent diagram

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// collectMissingDiagrams returns diagnostics for @state sequences that reference
// diagrams not present in diagramByID.
func collectMissingDiagrams(fn ssac.ServiceFunc, diagramByID map[string]*statemachine.StateDiagram) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, seq := range fn.Sequences {
		if seq.Type != "state" {
			continue
		}
		if _, ok := diagramByID[seq.DiagramID]; ok {
			continue
		}
		file := fn.FileName
		if file == "" {
			file = "ssac/" + fn.Name + ".ssac"
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        file,
			Line:        seq.Line,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     "[XMS-24] @state references diagram \"" + seq.DiagramID + "\" which does not exist",
			Advice:      "Add the file specs/states/" + seq.DiagramID + ".mmd",
			OperationID: fn.Name,
		})
	}
	return diags
}
