//ff:func feature=validate type=util control=iteration dimension=1 topic=states
//ff:what collectInvalidTransitions — collects @state transitions in SSaC functions that do not match any diagram event

package ssac_statemachine

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
)

// collectInvalidTransitions returns diagnostics for @state sequences whose
// transition is not a valid event in the referenced diagram.
func collectInvalidTransitions(fn ssac.ServiceFunc, diagramByID map[string]*statemachine.StateDiagram) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, seq := range fn.Sequences {
		if d, ok := checkTransitionValidity(fn, seq, diagramByID); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
