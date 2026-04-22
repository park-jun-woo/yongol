//ff:func feature=validate type=rule control=sequence topic=ssac-sqlc
//ff:what Run — execute all SSaC↔sqlc cross-validation rules (XQS-*)
package ssac_sqlc

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all SSaC↔sqlc cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xqs14InputKeyCase(fs)...)
	diags = append(diags, xqs15InputKeyInitialism(fs)...)
	diags = append(diags, xqs16InputKeyMissing(fs)...)
	diags = append(diags, xqs17ParamKeyMissing(fs)...)
	diags = append(diags, xqs18ParamTypeMismatch(fs)...)
	return diags
}
