//ff:func feature=validate type=rule control=iteration dimension=1 topic=config-check
//ff:what XNO-50 — OpenAPI securityScheme → Manifest middleware 매칭 검사

package openapi_manifest

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xno50SecuritySchemeMiddleware validates XNO-50: each OpenAPI securityScheme
// must have a matching middleware entry in manifest.yaml backend.middleware.
func xno50SecuritySchemeMiddleware(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Components == nil || fs.Manifest == nil {
		return nil
	}
	mwSet := middlewareSet(fs.Manifest.Backend.Middleware)
	var diags []diagnostic.Diagnostic
	for name := range fs.OpenAPIDoc.Components.SecuritySchemes {
		if !mwSet[name] {
			diags = append(diags, diagnostic.Diagnostic{
				File:    "api/openapi.yaml",
				Phase:   diagnostic.PhaseValidate,
				Level:   diagnostic.LevelError,
				Message: "[XNO-50] OpenAPI securityScheme \"" + name + "\" has no matching middleware in manifest.yaml",
				Advice:  "manifest backend.middleware 에 \"" + name + "\" 를 추가하세요",
			})
		}
	}
	return diags
}
