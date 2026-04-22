//ff:func feature=validate type=rule control=iteration dimension=1 topic=openapi-structural
//ff:what O-1 — OpenAPI path에 같은 이름의 {param}이 중복 등장하는지 감지

package openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// o01PathParamConflict flags paths where the same parameter name appears
// in more than one segment (e.g. /users/{id}/posts/{id}).
func o01PathParamConflict(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for path := range fs.OpenAPIDoc.Paths.Map() {
		if hasPathParamConflict(path) {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "api/openapi.yaml",
				Line:    fs.OpenAPILines.PathLine(path),
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[O-1] path parameter conflict at " + path,
				Advice:  "동일 path 의 파라미터 이름을 일치시키세요",
			})
		}
	}
	return diags
}
