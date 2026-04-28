//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-03 — hurl 요청 body JSON 필드가 OpenAPI request schema 에 존재

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh03RequestFieldInSchema enforces XOH-03: every top-level JSON key
// in a hurl request body must appear in the OpenAPI request schema's
// properties. Typos like `emale` vs `email` surface immediately.
//
// Scope limited to JSON object bodies — arrays / primitives / multipart
// forms are ignored (parsed BodyFields is empty in those cases). When
// the matched operation declares no requestBody, the rule skips the
// entry: an extra payload is a rule elsewhere (or a GET with body).
func xoh03RequestFieldInSchema(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		diags = append(diags, xoh03CheckEntry(e, routes)...)
	}
	return diags
}
