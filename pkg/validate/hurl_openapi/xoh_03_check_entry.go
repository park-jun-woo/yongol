//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what xoh03CheckEntry — 한 hurl entry 의 body 필드 ↔ OpenAPI request schema 비교

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// xoh03CheckEntry returns XOH-03 diagnostics for one hurl entry, or nil
// when the entry lacks body fields / lacks a matched operation / the
// operation declares no usable JSON requestBody.
func xoh03CheckEntry(e hurl.HurlEntry, routes []apiRoute) []diagnostic.Diagnostic {
	if len(e.BodyFields) == 0 {
		return nil
	}
	segs := normalizeHurlPath(e.Path)
	route := findExactRoute(segs, e.Method, routes)
	if route == nil || route.Op == nil {
		return nil
	}
	props, ok := requestBodyProps(route.Op)
	if !ok {
		return nil
	}
	available := sortedKeys(boolSet(props))
	var diags []diagnostic.Diagnostic
	for _, f := range e.BodyFields {
		if _, ok := props[f]; ok {
			continue
		}
		diags = append(diags, diagnostic.Diagnostic{
			File:  e.File,
			Line:  e.Line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: "[XOH-03] field \"" + f + "\" absent from " + opLabel(route.Op) +
				" requestBody schema",
			Advice: "available fields: " + joinKeys(available),
		})
	}
	return diags
}
