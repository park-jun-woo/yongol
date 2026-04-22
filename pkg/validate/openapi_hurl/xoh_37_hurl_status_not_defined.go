//ff:func feature=validate type=rule control=iteration dimension=1 topic=scenario-check
//ff:what XOH-37 — verifies that a Hurl entry's expected status code is defined in the OpenAPI responses

package openapi_hurl

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh37HurlStatusNotDefined validates XOH-37: when a hurl entry exactly
// matches an OpenAPI operation, its expected status code must be listed in
// that operation's responses. Path/method mismatches are handled by XOH-35/36.
func xoh37HurlStatusNotDefined(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		if e.StatusCode == "" {
			continue
		}
		segs := normalizeHurlPath(e.Path)
		route := findExactRoute(segs, e.Method, routes)
		if route == nil {
			continue
		}
		if route.Responses[e.StatusCode] {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:    e.File,
			Line:    e.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelWarning,
			Message: "[XOH-37] status " + e.StatusCode + " for " + e.Method + " " + e.Path + " not defined in OpenAPI responses",
			Advice:  "Add " + e.StatusCode + " to the OpenAPI op responses, or change the expected status code in Hurl",
		})
	}
	return diags
}
