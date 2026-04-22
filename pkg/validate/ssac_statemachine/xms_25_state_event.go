//ff:func feature=validate type=rule control=iteration dimension=1 topic=states
//ff:what XMS-25 — @state transition → diagram event

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xms25StateEvent validates XMS-25: @state transition → diagram event
func xms25StateEvent(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	diagramByID := buildDiagramByID(fs.StateDiagrams)
	var diags []diagnostic.Diagnostic
	for _, fn := range fs.ServiceFuncs {
		diags = append(diags, collectInvalidTransitions(fn, diagramByID)...)
	}
	return diags
}
