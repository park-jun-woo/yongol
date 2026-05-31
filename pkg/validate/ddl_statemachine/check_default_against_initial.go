//ff:func feature=validate type=util control=sequence topic=ddl-statemachine
//ff:what checkDefaultAgainstInitial — compares the initial state of a single diagram against the DDL DEFAULT value

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
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: "[XDM-28] DDL " + table + "." + column + " has no DEFAULT but stateDiagram " + diagramID +
				" initial state is '" + initial + "'",
			Advice: "Add DEFAULT '" + initial + "' to DDL column " + table + "." + column,
		}
	}
	if got != initial {
		return &diagnostic.Diagnostic{
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: "[XDM-28] DDL " + table + "." + column + " DEFAULT '" + got + "' ≠ stateDiagram " +
				diagramID + " initial '" + initial + "'",
			Advice: "Change the DDL DEFAULT to '" + initial + "' or change the stateDiagram [*] --> transition to '" + got + "' so they agree",
		}
	}
	return nil
}
