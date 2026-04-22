//ff:func feature=validate type=rule control=iteration dimension=1 topic=ddl-statemachine
//ff:what XDM-27 — @state field → DDL column

package ddl_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xdm27StateFieldColumn validates XDM-27: Mermaid stateDiagram DiagramID must
// correspond to an existing DDL column. Detail logic delegated to
// checkStateFieldColumn for one diagram at a time.
func xdm27StateFieldColumn(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	g := fs.Ground()
	if g == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, d := range fs.StateDiagrams {
		if d == nil {
			continue
		}
		if diag := checkStateFieldColumn(d.ID, g); diag != nil {
			diags = append(diags, *diag)
		}
	}
	return diags
}
