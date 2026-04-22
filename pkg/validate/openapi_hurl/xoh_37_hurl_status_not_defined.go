//ff:func feature=validate type=rule control=iteration dimension=1 topic=scenario-check
//ff:what XOH-37 — Hurl entry의 기대 상태코드가 OpenAPI responses에 정의되어 있는지 검증

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
			Advice:  "OpenAPI op responses 에 " + e.StatusCode + " 를 추가하거나 Hurl 에서 기대 상태코드를 변경하세요",
		})
	}
	return diags
}
