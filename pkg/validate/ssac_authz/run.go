//ff:func feature=validate type=rule control=sequence topic=ssac-authz
//ff:what Run — SSaC↔Authz 교차 검증 실행 (XAS-*)
package ssac_authz

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all SSaC↔Authz cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xas60AuthInputField(fs)...)
	return diags
}
