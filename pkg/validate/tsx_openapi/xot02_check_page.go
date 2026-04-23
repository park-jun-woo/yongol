//ff:func feature=validate type=util control=iteration dimension=1 topic=tsx-openapi
//ff:what XOT-2 헬퍼 — 단일 TSX 페이지의 모든 apiClient 호출 인자를 검사

package tsx_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	tsxparser "github.com/park-jun-woo/yongol/pkg/parser/tsx"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

// xot02CheckPage evaluates XOT-2 for every apiClient call on a single page.
// Delegates per-call argument checks to xot02CheckCall to stay under the Q1
// depth budget.
func xot02CheckPage(page tsxparser.PageSpec, opIDs rule.StringSet, lookup map[string]rule.StringSet) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, call := range page.Calls {
		if !opIDs[call.OperationID] {
			continue // XOT-1 already reports this
		}
		params := lookup["OpenAPI.param."+call.OperationID]
		diags = append(diags, xot02CheckCall(page.File, call, params)...)
	}
	return diags
}
