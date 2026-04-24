//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-02 — hurl `HTTP <status>` 가 OpenAPI responses 에 선언됨

package hurl_openapi

import (
	"sort"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh02StatusDeclared enforces XOH-02: when a hurl entry matches an
// OpenAPI operation exactly, the asserted HTTP status must be listed in
// that operation's responses. Entries without a status or without a
// matched operation are skipped — XOH-01 already reports on the latter.
//
// This rule replaces the former XOH-37. Raising severity from WARNING
// to ERROR matches the Phase002 plan: drift between hurl and OpenAPI
// responses is an SSOT disagreement that blocks generate.
func xoh02StatusDeclared(fs *yongol.Fullstack) []diagnostic.Diagnostic {
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
		available := sortedKeys(route.Responses)
		diags = append(diags, diagnostic.Diagnostic{
			File:  e.File,
			Line:  e.Line,
			Phase: diagnostic.PhaseValidate,
			Level: diagnostic.LevelError,
			Message: "[XOH-02] status " + e.StatusCode + " for " + e.Method + " " + e.Path +
				" — not declared in OpenAPI responses " + joinKeys(available),
			Advice: "Add " + e.StatusCode + " to the operation's responses, or change the hurl HTTP status",
		})
	}
	return diags
}

// sortedKeys returns the keys of a bool map in deterministic order so
// diagnostics are stable in tests and CI diff.
func sortedKeys(m map[string]bool) []string {
	out := make([]string, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	sort.Strings(out)
	return out
}

// joinKeys renders a key list as `[a, b, c]` for use in diagnostic
// messages.
func joinKeys(keys []string) string {
	if len(keys) == 0 {
		return "[]"
	}
	return "[" + joinCSV(keys) + "]"
}

// joinCSV is a trivial comma-separated join, kept local so we do not
// reach for strings.Join where it would pull in the `strings` import
// only for this one-line use.
func joinCSV(s []string) string {
	out := ""
	for i, v := range s {
		if i > 0 {
			out += ", "
		}
		out += v
	}
	return out
}
