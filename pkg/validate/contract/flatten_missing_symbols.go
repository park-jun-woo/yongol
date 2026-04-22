//ff:func feature=validate-contract type=util control=iteration dimension=1
//ff:what flattenMissingSymbols — missingSymbols 3 범주를 Diagnostic 슬라이스로 직렬화

package contract

import "github.com/park-jun-woo/yongol/pkg/diagnostic"

// flattenMissingSymbols emits one Diagnostic per missing entry, in
// category order (queries → calls → fields). Keeping the emission
// routine separate from the per-file inspector lets checkOne* stay a
// simple composition of narrower helpers.
func flattenMissingSymbols(path string, ms missingSymbols) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	for _, q := range ms.Queries {
		diags = append(diags, diagnoseMissingQuery(path, q))
	}
	for _, c := range ms.Calls {
		diags = append(diags, diagnoseMissingCall(path, c))
	}
	for _, f := range ms.Fields {
		diags = append(diags, diagnoseMissingField(path, f))
	}
	return diags
}
