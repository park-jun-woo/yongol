//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh04CheckEntry — 한 hurl entry 의 [Asserts] jsonpath 를 schema 로 검증

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// xoh04CheckEntry returns XOH-04 diagnostics for the asserts declared
// on one hurl entry. Entries without asserts / without a matched
// operation / without a usable response schema contribute nothing.
func xoh04CheckEntry(e hurl.HurlEntry, routes []apiRoute) []diagnostic.Diagnostic {
	if len(e.Asserts) == 0 {
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
	for _, a := range e.Asserts {
		if jsonPathReachable(a.JSONPath, schema) {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:  e.File,
			Line:  a.Line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: "[XOH-04] jsonpath \"" + a.JSONPath + "\" not reachable in " +
				opLabel(route.Op) + " response schema",
			Advice: "Check the field name in openapi.yaml responses, or update the hurl assertion",
		})
	}
	return diags
}
