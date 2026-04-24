//ff:func feature=validate type=util control=sequence topic=states
//ff:what matchStatefulDiagram — 단일 diagram 이 주어진 resource 와 stateful 대응되는지 확인

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// matchStatefulDiagram returns a *statefulTarget when d describes the
// given resource (plural table name match) AND its `[*] --> X` initial
// state matches the DDL DEFAULT on the corresponding (table, column)
// via the Ground lookup. Returns nil on any precondition miss. Extracted
// from isStatefulResource so the outer func stays at iteration
// dimension=1 (single loop at depth 1).
func matchStatefulDiagram(d *statemachine.StateDiagram, plural, singular string, g *rule.Ground) *statefulTarget {
	if d == nil {
		return nil
	}
	table, column := diagramIDToTable(d.ID)
	if table != plural {
		return nil
	}
	if d.InitialState == "" {
		return nil
	}
	if g == nil {
		return nil
	}
	got := g.Types["DDL.default.value."+table+"."+column]
	if got == "" || got != d.InitialState {
		return nil
	}
	return &statefulTarget{
		Resource:    singular,
		Table:       table,
		Diagram:     d,
		StateColumn: column,
		Model:       pascalCaseFromLower(singular),
	}
}
