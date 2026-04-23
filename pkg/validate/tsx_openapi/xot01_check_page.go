//ff:func feature=validate type=util control=iteration dimension=1 topic=tsx-openapi
//ff:what XOT-1 헬퍼 — 단일 TSX PageSpec 안의 apiClient 호출을 검사

package tsx_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	tsxparser "github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// xot01CheckPage scans a single TSX PageSpec for apiClient calls whose
// operationId is missing from the OpenAPI lookup set.
func xot01CheckPage(page tsxparser.PageSpec, opIDs rule.StringSet) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, call := range page.Calls {
		if opIDs[call.OperationID] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    page.File,
			Line:    call.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XOT-1] apiClient." + call.OperationID + "() has no matching OpenAPI operationId",
			Advice:  "Add operationId: " + call.OperationID + " to openapi.yaml, or check for a typo in the call name",
		})
	}
	return diags
}
