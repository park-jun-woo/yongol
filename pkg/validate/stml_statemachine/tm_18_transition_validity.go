//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-statemachine
//ff:what TM-18 — data-action 전이가 data-enabled-when이 건 상태에서 stateDiagram상 합법인지 검사 (WARNING)

package stml_statemachine

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm18TransitionValidity checks that the transition named by an action's
// data-action (OperationID) is legal from the state its data-enabled-when guard
// requires. The guard model prefix selects the diagram (modelSymbol →
// diagramBySymbol); a model that matches no diagram is a no-op. ValidFromStates
// gives the legal source states for that event; if the guarded state is not
// among them, the action would enable an illegal transition (WARNING — design
// allows intentional exceptions). Guard syntax errors are reported by TM-17, so
// a parse failure here is silently skipped.
func tm18TransitionValidity(a stml.ActionBlock, file string, diagramBySymbol map[string]*statemachine.StateDiagram) []diagnostic.Diagnostic {
	if a.EnabledWhen == "" {
		return nil
	}
	expr, err := stml.ParseGuard(a.EnabledWhen)
	if err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, pair := range collectComparePairs(expr) {
		d, ok := diagramBySymbol[modelSymbol(pair.Model)]
		if !ok {
			continue
		}
		if stateInSlice(d.ValidFromStates(a.OperationID), pair.Value) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelWarning,
			Message:     fmt.Sprintf("[TM-18] action %q is enabled when %q is %q, but that transition is not legal from state %q in stateDiagram %q", a.OperationID, pair.Model, pair.Value, pair.Value, modelSymbol(pair.Model)),
			Advice:      fmt.Sprintf("Change data-enabled-when to a state from which %q is a legal transition, or add that transition to the %q stateDiagram", a.OperationID, modelSymbol(pair.Model)),
			OperationID: a.OperationID,
		})
	}
	return diags
}
