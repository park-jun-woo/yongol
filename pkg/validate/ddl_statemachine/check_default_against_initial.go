//ff:func feature=validate type=util control=sequence topic=ddl-statemachine
//ff:what checkDefaultAgainstInitial — 단일 diagram 의 initial 과 DDL DEFAULT 대조

package ddl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// checkDefaultAgainstInitial returns an XDM-28 ERROR when the DDL DEFAULT on
// `<table>.<column>` is absent or does not match `initial`. Returns nil on
// match or when the target DDL column has no registered DEFAULT (treat as
// "not declared" → ERROR).
func checkDefaultAgainstInitial(g *rule.Ground, diagramID, table, column, initial string) *diagnostic.Diagnostic {
	got := g.Types["DDL.default.value."+table+"."+column]
	if got == "" {
		return &diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XDM-28] DDL " + table + "." + column + " has no DEFAULT but stateDiagram " + diagramID +
				" initial state is '" + initial + "'",
			Advice: "DDL " + table + "." + column + " 에 DEFAULT '" + initial + "' 를 추가하세요",
		}
	}
	if got != initial {
		return &diagnostic.Diagnostic{
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XDM-28] DDL " + table + "." + column + " DEFAULT '" + got + "' ≠ stateDiagram " +
				diagramID + " initial '" + initial + "'",
			Advice: "DDL DEFAULT 를 '" + initial + "' 로 변경하거나 stateDiagram [*] --> 를 '" + got + "' 로 변경하여 통일하세요",
		}
	}
	return nil
}
