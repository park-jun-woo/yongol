//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what SSaC 함수의 @state sequence 중 존재하지 않는 diagram 수집

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
			File:    file,
			Line:    seq.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XMS-24] @state references diagram \"" + seq.DiagramID + "\" which does not exist",
			Advice:  "specs/states/" + seq.DiagramID + ".mmd 파일을 추가하세요",
		})
	}
	return diags
}
