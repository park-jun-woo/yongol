//ff:func feature=validate type=rule control=sequence topic=openapi-hurl
//ff:what Run — OpenAPI↔Hurl 교차 검증 실행 (XOH-*)
package openapi_hurl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all OpenAPI↔Hurl cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xoh35HurlPathOpenAPI(fs)...)
	diags = append(diags, xoh36HurlMethodOpenAPI(fs)...)
	diags = append(diags, xoh37HurlStatusNotDefined(fs)...)
	return diags
}
