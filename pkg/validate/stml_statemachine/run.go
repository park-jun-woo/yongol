//ff:func feature=validate type=rule control=iteration dimension=2 topic=stml-statemachine
//ff:what Run — STML↔stateDiagram 교차 검증 실행 (TM-15, TM-18, TM-23). 한쪽이라도 없으면 no-op.

package stml_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all STML↔stateDiagram cross-validation rules (TM-15, TM-18,
// TM-23). It is a no-op unless both STML pages and stateDiagrams are present,
// so projects without a stateDiagram see no change in validation output.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if len(fs.STMLPages) == 0 || len(fs.StateDiagrams) == 0 {
		return nil
	}

	stateMap := buildDiagramStateMap(fs.StateDiagrams)
	diagramBySymbol := buildDiagramBySymbol(fs.StateDiagrams)

	var diags []diagnostic.Diagnostic
	for _, page := range fs.STMLPages {
		for _, a := range page.Actions {
			diags = append(diags, tm15StateValueInDiagram(a, page.FileName, stateMap)...)
			diags = append(diags, tm18TransitionValidity(a, page.FileName, diagramBySymbol)...)
			diags = append(diags, tm23RedirectStateConflict(a, page.FileName, diagramBySymbol, fs.STMLPages)...)
		}
	}
	return diags
}
