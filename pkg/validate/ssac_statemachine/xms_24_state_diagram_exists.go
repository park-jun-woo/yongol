//ff:func feature=validate type=rule control=iteration dimension=1 topic=states
//ff:what XMS-24 — verifies that every @state references an existing diagram

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xms24StateDiagramExists validates XMS-24: every @state must reference an existing diagram.
func xms24StateDiagramExists(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	diagramByID := buildDiagramByID(fs.StateDiagrams)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		diags = append(diags, collectMissingDiagrams(fn, diagramByID)...)
	}
	return diags
}
