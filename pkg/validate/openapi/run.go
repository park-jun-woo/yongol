//ff:func feature=validate type=rule control=sequence topic=openapi-structural
//ff:what Run — OpenAPI 검증 전체 실행 (O-*, XOO-*)
package openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all OpenAPI validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, o01PathParamConflict(fs)...)
	diags = append(diags, o02PathParamCaseConflict(fs)...)
	diags = append(diags, o03PathTemplateParam(fs)...)
	diags = append(diags, o04OpIdRequired(fs)...)
	diags = append(diags, o05ResponseBodyRequired(fs)...)
	diags = append(diags, o06RequiredPropertyConsistency(fs)...)
	diags = append(diags, xoo71PasswordNoMinLength(fs)...)
	diags = append(diags, xoo72EmailNoFormat(fs)...)
	diags = append(diags, xoe01ErrorResponseRequired(fs)...)
	return diags
}
