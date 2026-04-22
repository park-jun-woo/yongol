//ff:func feature=validate type=rule control=iteration dimension=1 topic=scenario-check
//ff:what XOH-36 — verifies that a Hurl entry's HTTP method is defined in the matched OpenAPI path

package openapi_hurl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh36HurlMethodOpenAPI validates XOH-36: when a hurl entry's path matches an
// OpenAPI path, its HTTP method must also be defined on that path. Missing
// paths are handled by XOH-35 and skipped here to avoid duplicate reports.
func xoh36HurlMethodOpenAPI(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		pathIdx, route := matchHurlEntry(e, routes)
		if pathIdx < 0 || route != nil {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    e.File,
			Line:    e.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: "[XOH-36] method " + e.Method + " " + e.Path + " — path exists but method not defined in OpenAPI",
			Advice:  "Align the Hurl method with the OpenAPI operation's method, or add the method to OpenAPI",
		})
	}
	return diags
}
