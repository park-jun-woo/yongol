//ff:func feature=validate type=rule control=iteration dimension=1 topic=openapi-structural
//ff:what O-1 — detects duplicate {param} names in an OpenAPI path

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
				Advice:  "Use distinct parameter names within the same path",
			})
		}
	}
	return diags
}
