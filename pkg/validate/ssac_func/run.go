//ff:func feature=validate type=rule control=sequence topic=ssac-func
//ff:what Run — execute all SSaC↔FuncSpec cross-validation rules (XFS-*, XSF-*)
package ssac_func

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all SSaC↔FuncSpec cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xfs39CallToFuncSpec(fs)...)
	diags = append(diags, xfs42CallInputsCount(fs)...)
	diags = append(diags, xfs43CallInputFields(fs)...)
	diags = append(diags, xfs44CallInputType(fs)...)
	diags = append(diags, xfs45CallResultMissing(fs)...)
	diags = append(diags, xsf46CallResultIgnored(fs)...)
	diags = append(diags, xfs63CallFuncSignature(fs)...)
	diags = append(diags, xsf62FuncSpecUsed(fs)...)
	diags = append(diags, xfs70AuthInputType(fs)...)
	diags = append(diags, xfs73CallRequestParamType(fs)...)
	return diags
}
