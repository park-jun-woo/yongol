//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-08 — hurl [Captures] 의 jsonpath 가 OpenAPI response schema 에 존재

package hurl_openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// xoh08CapturePathInSchema enforces XOH-08: every [Captures] line that
// uses `jsonpath "$..."` must refer to a path present in the operation's
// response schema. Without this rule, a capture against a non-existent
// field silently returns empty and only fails at runtime when a later
// step dereferences the variable (exactly the failure mode in BUG-031).
func xoh08CapturePathInSchema(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs == nil || fs.OpenAPIDoc == nil {
		return nil
	}
	routes := collectOpenAPIRoutes(fs.OpenAPIDoc)
	var diags []diagnostic.Diagnostic
	for _, e := range fs.HurlEntries {
		diags = append(diags, xoh08CheckEntry(e, routes)...)
	}
	return diags
}
