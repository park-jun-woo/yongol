//ff:func feature=validate type=util control=sequence topic=policy-check
//ff:what xsp29MissingSSaCDiag — Rego pair 가 SSaC 에 없을 때 XSP-29 진단 생성

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
		Advice: fmt.Sprintf("SSaC 함수에 @auth Action=\"%s\" Resource=\"%s\" 시퀀스를 추가하세요", pair[0], pair[1]),
	}, true
}
