//ff:func feature=validate type=rule control=sequence topic=hurl-openapi
//ff:what checkEntryURLMethod — 단일 hurl entry 의 URL+method 가 OpenAPI route 와 일치하는지 검사

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/parser/hurl"
)

// checkEntryURLMethod evaluates one hurl entry against the OpenAPI routes
// for XOH-01. It returns nil when the entry is an external-service call
// ({{authurl}}/{{rest}} etc.) or matches a declared operation; otherwise it
// returns the drift diagnostic. Entries using {{host}} or an absolute
// http(s):// URL (URLVar == "") are the user's own API and stay matched.
func checkEntryURLMethod(e hurl.HurlEntry, routes []apiRoute) *diagnostic.Diagnostic {
	if e.URLVar != "" && e.URLVar != "host" {
		return nil
	}
	segs := normalizeHurlPath(e.Path)
	if findExactRoute(segs, e.Method, routes) != nil {
		return nil
	}
	msg, advice := xoh01Message(e, segs, routes)
	return &diagnostic.Diagnostic{
		File:    e.File,
		Line:    e.Line,
		Phase:   diagnostic.PhaseValidate,
		Level:   diagnostic.LevelError,
		Message: msg,
		Advice:  advice,
	}
}
