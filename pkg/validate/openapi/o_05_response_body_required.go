//ff:func feature=validate type=rule control=iteration dimension=2 topic=response-body-required
//ff:what O-5 — every 4xx/5xx response must declare content: application/json + schema

package openapi

import (
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// o05ResponseBodyRequired validates O-5: every operation response with a
// status code in the 4xx or 5xx range must declare `content: application/json`
// with a schema (inline or $ref). 204 No Content and 304 Not Modified are
// exempt because they are intentionally bodyless. 1xx/2xx/3xx codes fall
// outside this rule's scope.
//
// Rationale: yongol = SaaS / business backend orchestrator. Every error
// response must carry a structured body so frontends can render error
// messages, clients can distinguish error causes, and logs/alerts capture
// machine-readable codes. Empty 4xx/5xx is anti-pattern in yongol's target
// domain (RFC 7807 Problem Details recommended, but not enforced).
//
// Side effect: oapi-codegen emits `<Op><Status>JSONResponse` only when the
// response declares JSON content with a schema. Without this rule, a missing
// content block silently degrades the type to `<Op><Status>Response`
// (struct{}), causing yongol's SSaC handler emit (which always references
// `JSONResponse`) to fail compilation. O-5 closes that gap structurally.
func o05ResponseBodyRequired(fs *yongol.Fullstack) []diagnostic.Diagnostic {
	if fs.OpenAPIDoc == nil || fs.OpenAPIDoc.Paths == nil {
		return nil
	}
	var diags []diagnostic.Diagnostic
	for path, item := range fs.OpenAPIDoc.Paths.Map() {
		if item == nil {
			continue
		}
		for _, op := range item.Operations() {
			diags = append(diags, checkOpResponseBodies(op, path, fs.OpenAPILines)...)
		}
	}
	return diags
}
