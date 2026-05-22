//ff:func feature=validate type=util control=sequence topic=policy-check
//ff:what xps28MissingRegoDiag — generates an XPS-28 diagnostic when an SSaC pair is absent from the Rego policy

package ssac_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// xps28MissingRegoDiag returns (diag, true) when the SSaC (action, resource)
// pair has no matching Rego allow rule.
func xps28MissingRegoDiag(pair [2]string, regoPairs map[[2]string]bool, pairLoc map[[2]string]PairLocation) (diagnostic.Diagnostic, bool) {
	if regoPairs[pair] {
		return diagnostic.Diagnostic{}, false
	}
	loc := pairLoc[pair]
	return diagnostic.Diagnostic{
		File:        loc.File,
		Line:        loc.Line,
		Phase:       diagnostic.PhaseValidate,
		Level:       diagnostic.LevelError,
		OperationID: pair[0],
		Message: fmt.Sprintf(
			"[XPS-28] SSaC authorize (%s, %s) has no matching Rego allow rule",
			pair[0], pair[1]),
		Advice: fmt.Sprintf("Add an allow rule for (action: %s, resource: %s) to the Rego policy", pair[0], pair[1]),
	}, true
}
