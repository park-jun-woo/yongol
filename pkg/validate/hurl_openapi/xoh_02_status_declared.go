//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-02 — hurl `HTTP <status>` 가 OpenAPI responses 에 선언됨

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh02StatusDeclared enforces XOH-02: when a hurl entry matches an
// OpenAPI operation exactly, the asserted HTTP status must be listed in
// that operation's responses. Entries without a status or without a
// matched operation are skipped — XOH-01 already reports on the latter.
func xoh02StatusDeclared(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		if d, ok := xoh02CheckEntry(e, routes); ok {
			diags = append(diags, d)
		}
	}
	return diags
}
