//ff:func feature=validate type=rule control=iteration dimension=1 topic=states
//ff:what XSM-23 — transition event → SSaC

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xsm23TransitionToFunc validates XSM-23: transition event → SSaC
func xsm23TransitionToFunc(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	funcNames := buildFuncNameSet(fs.ServiceFuncs)
	var diags []diagnostic.Diagnostic
	for _, d := range fs.StateDiagrams {
		diags = append(diags, collectUnmatchedEvents(d, funcNames)...)
	}
	return diags
}
