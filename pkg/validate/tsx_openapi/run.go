//ff:func feature=validate type=rule control=sequence topic=tsx-openapi
//ff:what Run — TSX ↔ OpenAPI 교차 검증 (XOT-*) 실행 진입점
package tsx_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all TSX → OpenAPI cross-validation rules.
// Direction: TSX (claim) → OpenAPI (truth). Reverse direction (unused
// operations) is intentionally not enforced — see plans/tsx/Phase002.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xot01OperationID(fs)...)
	diags = append(diags, xot02ParameterMatch(fs)...)
	diags = append(diags, xot03FormField(fs)...)
	return diags
}
