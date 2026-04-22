//ff:func feature=validate type=util control=sequence topic=policy-check
//ff:what xsp29MissingSSaCDiag — generates an XSP-29 diagnostic when a Rego pair is absent from SSaC

package ssac_rego

import (
	"fmt"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// xsp29MissingSSaCDiag returns (diag, true) when the Rego (action, resource)
// pair has no matching SSaC @auth sequence.
func xsp29MissingSSaCDiag(pair [2]string, ssacPairs map[[2]string]bool, pairLoc map[[2]string]PairLocation) (diagnostic.Diagnostic, bool) {
	if ssacPairs[pair] {
		return diagnostic.Diagnostic{}, false
	}
	loc := pairLoc[pair]
	return diagnostic.Diagnostic{
		File:  loc.File,
		Line:  loc.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XSP-29] Rego allow rule (%s, %s) has no matching SSaC @auth sequence",
			pair[0], pair[1]),
		Advice: fmt.Sprintf("Add an @auth Action=\"%s\" Resource=\"%s\" sequence to the SSaC function", pair[0], pair[1]),
	}, true
}
