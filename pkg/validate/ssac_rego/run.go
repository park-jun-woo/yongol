//ff:func feature=validate type=rule control=sequence topic=ssac-rego
//ff:what Run — SSaC↔Rego 교차 검증 실행 (XPS-*, XSP-*)
package ssac_rego

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all SSaC↔Rego cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xps28SSaCAuthToRego(fs)...)
	diags = append(diags, xsp29RegoAllowToSSaC(fs)...)
	return diags
}
