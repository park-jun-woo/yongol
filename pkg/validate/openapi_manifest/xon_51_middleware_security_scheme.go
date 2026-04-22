//ff:func feature=validate type=rule control=iteration dimension=1 topic=config-check
//ff:what XON-51 — Manifest middleware → OpenAPI securityScheme 매칭 검사

package openapi_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xon51MiddlewareSecurityScheme validates XON-51: each middleware in
// manifest.yaml backend.middleware must have a matching OpenAPI securityScheme.
func xon51MiddlewareSecurityScheme(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.Manifest == nil {
		return nil
	}
	schemes := schemeNameSet(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, m := range fs.Manifest.Backend.Middleware {
		if !schemes[m] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "manifest.yaml",
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[XON-51] middleware \"" + m + "\" has no matching OpenAPI securityScheme",
				Advice:  "OpenAPI components.securitySchemes 에 \"" + m + "\" 를 추가하세요",
			})
		}
	}
	return diags
}
