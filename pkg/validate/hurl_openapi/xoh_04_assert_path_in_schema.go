//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-04 — hurl [Asserts] jsonpath 가 OpenAPI response schema 에 도달 가능

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh04AssertPathInSchema enforces XOH-04: every `jsonpath "$.path"` in
// a hurl [Asserts] block must resolve to a field declared in the
// matched operation's response schema for the asserted status code.
// When the status-specific response lacks content, we fall back to the
// 2xx default so assertions like `jsonpath "$.id"` on the canonical
// happy path keep working even when responses are split by status.
func xoh04AssertPathInSchema(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		diags = append(diags, xoh04CheckEntry(e, routes)...)
	}
	return diags
}
