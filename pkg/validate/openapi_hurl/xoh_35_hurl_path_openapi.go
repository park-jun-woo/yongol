//ff:func feature=validate type=rule control=iteration dimension=1 topic=scenario-check
//ff:what XOH-35 — verifies that a Hurl path matches a path defined in OpenAPI

package openapi_hurl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh35HurlPathOpenAPI validates XOH-35: every Hurl entry path must match
// at least one OpenAPI path (ignoring method / status).
func xoh35HurlPathOpenAPI(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		if pathIdx, _ := matchHurlEntry(e, routes); pathIdx >= 0 {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    e.File,
			Line:    e.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XOH-35] hurl path " + e.Path + " not found in OpenAPI",
			Advice:  "Align the Hurl path with an OpenAPI path, or add the path to OpenAPI",
		})
	}
	return diags
}
