//ff:func feature=validate type=rule control=iteration dimension=2 topic=config-check
//ff:what XNO-52 — OpenAPI endpoint security 참조가 Manifest middleware에 존재하는지 검사

package openapi_manifest

import (
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xno52SecurityMiddleware validates XNO-52: every security requirement name
// referenced by an endpoint must exist in manifest.yaml backend.middleware.
func xno52SecurityMiddleware(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil || fs.Manifest == nil {
		return nil
	}
	mwSet := middlewareSet(fs.Manifest.Backend.Middleware)
	var diags []diagnostic.Diagnostic
	for pathStr, pathItem := range fs.OpenAPIDoc.Paths.Map() {
		for method, op := range pathItem.Operations() {
			line := fs.OpenAPILines.OperationLine(op.OperationID)
			if line == 0 {
				line = fs.OpenAPILines.PathLine(pathStr)
			}
			for _, name := range missingSecurityNames(op, mwSet) {
				diags = append(diags, diagnostic.Diagnostic{
					File:    "api/openapi.yaml",
					Line:    line,
					Phase:   diagnostic.PhaseValidate,
					Level:   diagnostic.LevelError,
					Message: "[XNO-52] " + strings.ToUpper(method) + " " + pathStr + " references security \"" + name + "\" not in manifest.yaml middleware",
					Advice:  "manifest backend.middleware 에 security \"" + name + "\" 의 핸들러를 추가하세요",
				})
			}
		}
	}
	return diags
}
