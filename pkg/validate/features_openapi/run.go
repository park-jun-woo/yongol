//ff:func feature=validate type=rule control=sequence topic=features-openapi
//ff:what Run — Features↔OpenAPI 교차 검증 실행 (XFO-*, XOF-*)
package features_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// Run executes all Features↔OpenAPI cross-validation rules.
func Run(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	var diags []diagnostic.Diagnostic
	diags = append(diags, xfo01OpNotInOpenAPI(fs)...)
	diags = append(diags, xof01OpIDNotInFeatures(fs)...)
	return diags
}
