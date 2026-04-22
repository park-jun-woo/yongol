//ff:func feature=validate type=rule control=iteration dimension=1 topic=scenario-check
//ff:what XOH-36 — Hurl entry의 HTTP method가 매칭된 OpenAPI path에 정의되어 있는지 검증

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
			Advice:  "Hurl method 를 OpenAPI op 의 메서드와 일치시키거나 OpenAPI 에 해당 메서드를 추가하세요",
		})
	}
	return diags
}
