//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-01 — hurl 요청의 URL path + method 가 OpenAPI operation 으로 선언됨

package hurl_openapi

import (
	"sort"
	"strings"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh01URLMethod enforces XOH-01: every hurl request must hit an
// OpenAPI operation that declares the same path *and* method.
//
// This rule folds together the former XOH-35 (path only) and XOH-36
// (method on matched path). Splitting them produced two diagnostics for
// a single typo; a unified judgement keeps the advice concise and lets
// the XOH-03/04/08 rules skip cleanly when there is no matched operation
// to compare against.
func xoh01URLMethod(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		segs := normalizeHurlPath(e.Path)
		if findExactRoute(segs, e.Method, routes) != nil {
			continue
		}
		msg, advice := xoh01Message(e, segs, routes)
		diags = append(diags, diagnostic.Diagnostic{
			File:    e.File,
			Line:    e.Line,
			Phase:   diagnostic.PhaseValidate,
			Level:   diagnostic.LevelError,
			Message: msg,
			Advice:  advice,
		})
	}
	return diags
}

// xoh01Message picks the right diagnostic text based on whether the
// path exists at all. Listing the available methods when the path
// matches turns a drift report into a near copy-pasteable fix.
func xoh01Message(e hurl.HurlEntry, segs []string, routes []apiRoute) (string, string) {
	if findPathMatch(segs, routes) < 0 {
		return "[XOH-01] " + e.Method + " " + e.Path + " — path not declared in OpenAPI",
			"Add a matching operation to openapi.yaml, or fix the hurl request path"
	}
	methods := methodsForPath(segs, routes)
	return "[XOH-01] " + e.Method + " " + e.Path + " — method not declared on this path (OpenAPI lists " + strings.Join(methods, ", ") + ")",
		"Use one of the declared methods or add " + e.Method + " to the operation"
}

// methodsForPath returns the sorted list of HTTP methods declared for
// the first route whose segments match.
func methodsForPath(segs []string, routes []apiRoute) []string {
	var out []string
	for _, r := range routes {
		if segmentsMatch(segs, r.Segments) {
			out = append(out, r.Method)
		}
	}
	sort.Strings(out)
	return out
}
