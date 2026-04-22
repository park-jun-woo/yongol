//ff:func feature=validate type=rule control=iteration dimension=1 topic=scenario-check
//ff:what XOH-35 — Hurl path가 OpenAPI에 정의된 경로에 매칭되는지 검증

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
			Advice:  "Hurl path 를 OpenAPI 의 path 와 일치시키거나 OpenAPI 에 해당 경로를 추가하세요",
		})
	}
	return diags
}
