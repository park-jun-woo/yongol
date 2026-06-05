//ff:func feature=validate type=rule control=iteration dimension=1 topic=stml-statemachine
//ff:what TM-15 — 가드 비교식의 상태값(workflow.status=active의 active)이 참조 stateDiagram에 실존하는지 검사 (ERROR)

package stml_statemachine

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

// tm15StateValueInDiagram checks that every state value compared in an action's
// data-enabled-when guard exists in the matching stateDiagram. The guard model
// prefix is normalized to PascalCase (modelSymbol) and looked up in stateMap,
// which is keyed on StateDiagram.Symbol. A model that matches no diagram is not
// a TM-15 target (no-op). Guard syntax errors are reported earlier by TM-17, so
// a parse failure here is silently skipped.
func tm15StateValueInDiagram(a stml.ActionBlock, file string, stateMap map[string]map[string]bool) []diagnostic.Diagnostic {
	if a.EnabledWhen == "" {
		return nil
	}
	expr, err := stml.ParseGuard(a.EnabledWhen)
	if err != nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, pair := range collectComparePairs(expr) {
		states, ok := stateMap[modelSymbol(pair.Model)]
		if !ok {
			continue
		}
		if states[pair.Value] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:        file,
			Phase:       diagnostic.PhaseValidate,
			Level:       diagnostic.LevelError,
			Message:     fmt.Sprintf("[TM-15] data-enabled-when on action %q references state %q of %q, which is not a state in stateDiagram %q", a.OperationID, pair.Value, pair.Model, modelSymbol(pair.Model)),
			Advice:      fmt.Sprintf("Define state %q in the %q stateDiagram, or correct the guard value", pair.Value, modelSymbol(pair.Model)),
			OperationID: a.OperationID,
		})
	}
	return diags
}
