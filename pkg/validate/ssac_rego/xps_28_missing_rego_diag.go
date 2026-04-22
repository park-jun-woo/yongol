//ff:func feature=validate type=util control=sequence topic=policy-check
//ff:what xps28MissingRegoDiag — SSaC pair 가 Rego 에 없을 때 XPS-28 진단 생성

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
		File:  loc.File,
		Line:  loc.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf(
			"[XPS-28] SSaC authorize (%s, %s) has no matching Rego allow rule",
			pair[0], pair[1]),
		Advice: fmt.Sprintf("Rego policy 에 (action: %s, resource: %s) 의 allow 규칙을 추가하세요", pair[0], pair[1]),
	}, true
}
