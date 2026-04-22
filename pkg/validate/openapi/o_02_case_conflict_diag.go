//ff:func feature=validate type=util control=iteration dimension=1 topic=openapi-structural
//ff:what o02CaseConflictDiag — 단일 path parameter 의 case variants 를 진단으로 변환

package openapi

import (
	"fmt"
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

// o02CaseConflictDiag reports a [O-2] diagnostic when a single lowercase
// parameter name appears with more than one case variant across paths.
// Returns (Diagnostic, true) when a conflict exists; (zero, false) otherwise.
func o02CaseConflictDiag(lower string, cases map[string]map[string]bool) (diagnostic.Diagnostic, bool) {
	if len(cases) < 2 {
		return diagnostic.Diagnostic{}, false
	}
	// collect ordered variants for deterministic message
	variants := make([]string, 0, len(cases))
	for c := range cases {
		variants = append(variants, c)
	}
	sort.Strings(variants)
	return diagnostic.Diagnostic{
		File:  "api/openapi.yaml",
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: fmt.Sprintf("[O-2] path parameter %q appears with multiple case variants across paths: %v",
			lower, variants),
		Advice: "하나의 case 로 통일하세요 (case 구별은 router 에서 서로 다른 param 으로 처리됨)",
	}, true
}
