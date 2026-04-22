//ff:func feature=validate type=rule control=sequence topic=openapi-ssac
//ff:what Run — OpenAPI↔SSaC 교차 검증 실행 (XOS-*, XSO-*)
package openapi_ssac

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all OpenAPI↔SSaC cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xos15FuncNameOpID(fs)...)
	diags = append(diags, xso16OpIDToFunc(fs)...)
	diags = append(diags, xos17ResponseFields(fs)...)
	diags = append(diags, xso18ResponseFieldUsed(fs)...)
	diags = append(diags, xos19ShorthandResponse(fs)...)
	diags = append(diags, xso20ShorthandFieldUsed(fs)...)
	diags = append(diags, xos21ErrStatusNotInOpenAPI(fs)...)
	diags = append(diags, xos22ResponseNo2xx(fs)...)
	diags = append(diags, xos66UsedFieldsRequired(fs)...)
	diags = append(diags, xos67ResponseFieldType(fs)...)
	return diags
}
