//ff:func feature=validate type=rule control=sequence topic=hurl-openapi
//ff:what xoh02CheckEntry — 한 hurl entry 의 status 가 OpenAPI 에 선언됐는지 검사

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// xoh02CheckEntry returns the diagnostic for a single hurl entry when
// its status code is asserted but missing from the matched operation's
// responses. The second return signals whether a diagnostic was
// produced; false means the entry was skipped (no status, no matched
// route, or declared response).
func xoh02CheckEntry(e hurl.HurlEntry, routes []apiRoute) (diagnostic.Diagnostic, bool) {
	if e.StatusCode == "" {
		return diagnostic.Diagnostic{}, false
	}
	segs := normalizeHurlPath(e.Path)
	route := findExactRoute(segs, e.Method, routes)
	if route == nil {
		return diagnostic.Diagnostic{}, false
	}
	if route.Responses[e.StatusCode] {
		return diagnostic.Diagnostic{}, false
	}
	available := sortedKeys(route.Responses)
	return diagnostic.Diagnostic{
		File:  e.File,
		Line:  e.Line,
		Phase: diagnostic.PhaseValidate,
		Level: diagnostic.LevelError,
		Message: "[XOH-02] status " + e.StatusCode + " for " + e.Method + " " + e.Path +
			" — not declared in OpenAPI responses " + joinKeys(available),
		Advice: "Add " + e.StatusCode + " to the operation's responses, or change the hurl HTTP status",
	}, true
}
