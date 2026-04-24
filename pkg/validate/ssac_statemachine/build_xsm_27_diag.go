//ff:func feature=validate type=rule control=sequence topic=states
//ff:what buildXsm27Diag — XSM-27 WARNING 진단 조립 (Option A/B + self-loop 힌트)

package ssac_statemachine

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

// buildXsm27Diag assembles the full WARNING with the A/B options and the
// self-loop transition hint derived from the diagram's initial state. The
// `resultVar` is the variable name from the @get FindByID sequence; when
// empty the lowercase resource name is used as a readable fallback.
func buildXsm27Diag(fn ssac.ServiceFunc, target *statefulTarget, resultVar string) diagnostic.Diagnostic {
	file := fn.FileName
	if file == "" {
		file = "ssac/" + fn.Name + ".ssac"
	}
	diagramID := ""
	initial := ""
	if target.Diagram != nil {
		diagramID = target.Diagram.ID
		initial = target.Diagram.InitialState
	}
	msg := "[XSM-27] " + fn.Name + ": state-dependent operation on stateful resource '" +
		target.Resource + "' is missing @state declaration"

	varName := resultVar
	if varName == "" {
		varName = strings.ToLower(target.Resource)
	}

	advice := buildXsm27Advice(fn.Name, target, diagramID, initial, varName)

	return diagnostic.Diagnostic{
		File:    file,
		Line:    fn.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelWarning,
		Message: msg,
		Advice:  advice,
	}
}
