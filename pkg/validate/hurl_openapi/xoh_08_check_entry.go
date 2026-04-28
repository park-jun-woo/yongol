//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh08CheckEntry — 한 hurl entry 의 [Captures] jsonpath 를 schema 로 검증

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// xoh08CheckEntry returns XOH-08 diagnostics for captures declared on
// one hurl entry. Non-jsonpath sources / entries with no captures /
// entries whose operation lacks a usable response schema contribute
// nothing.
func xoh08CheckEntry(e hurl.HurlEntry, routes []apiRoute) []diagnostic.Diagnostic {
	if len(e.Captures) == 0 {
		return nil
	}
	segs := normalizeHurlPath(e.Path)
	route := findExactRoute(segs, e.Method, routes)
	if route == nil || route.Op == nil {
		return nil
	}
	schema := responseSchemaForStatus(route, e.StatusCode)
	if schema == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for _, c := range e.Captures {
		if c.Source != "jsonpath" || c.JSONPath == "" {
			continue
		}
		if jsonPathReachable(c.JSONPath, schema) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:  e.File,
			Line:  c.Line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: "[XOH-08] capture jsonpath \"" + c.JSONPath + "\" not in " +
				opLabel(route.Op) + " response — won't match at runtime",
			Advice: "Pick a field present in the OpenAPI response schema, or remove the capture",
		})
	}
	return diags
}
