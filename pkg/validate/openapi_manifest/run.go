//ff:func feature=validate type=rule control=sequence topic=openapi-manifest
//ff:what Run — OpenAPI↔Manifest 교차 검증 실행 (XNO-*, XON-*)
package openapi_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all OpenAPI↔Manifest cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xno50SecuritySchemeMiddleware(fs)...)
	diags = append(diags, xno52SecurityMiddleware(fs)...)
	diags = append(diags, xon51MiddlewareSecurityScheme(fs)...)
	diags = append(diags, xon60FrontendAuthTokenField(fs)...)
	diags = append(diags, sec04HTTPOverridesOperationID(fs)...)
	diags = append(diags, sec05RateLimitOpRoutable(fs)...)
	return diags
}
