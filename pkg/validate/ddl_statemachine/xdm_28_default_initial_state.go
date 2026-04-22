//ff:func feature=validate type=rule control=iteration dimension=1 topic=ddl-statemachine
//ff:what XDM-28 — stateDiagram [*] → X 초기 전이와 DDL DEFAULT 'X' 일치 검증

package ddl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdm28DefaultInitialState validates XDM-28: for each stateDiagram whose
// InitialState is declared (via `[*] --> X`), the corresponding DDL column
// must carry `DEFAULT '<X>'`. Target resolution lives in diagramDDLTarget;
// the mismatch check lives in checkDefaultAgainstInitial.
func xdm28DefaultInitialState(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, d := range fs.StateDiagrams {
		if d == nil || d.InitialState == "" {
			continue
		}
		table, column := diagramDDLTarget(d.ID)
		if diag := checkDefaultAgainstInitial(g, d.ID, table, column, d.InitialState); diag != nil {
			diags = append(diags, *diag)
		}
	}
	return diags
}
