//ff:func feature=validate type=rule control=iteration dimension=1 topic=hurl-openapi
//ff:what XOH-03 — hurl 요청 body JSON 필드가 OpenAPI request schema 에 존재

package hurl_openapi

import (
	"github.com/getkin/kin-openapi/openapi3"

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
		if len(e.BodyFields) == 0 {
			continue
		}
		segs := normalizeHurlPath(e.Path)
		route := findExactRoute(segs, e.Method, routes)
		if route == nil || route.Op == nil {
			continue
		}
		props, ok := requestBodyProps(route.Op)
		if !ok {
			continue
		}
		available := sortedKeys(boolSet(props))
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
	}
	return diags
}

// requestBodyProps returns the top-level properties of the first JSON
// content type declared on an operation's request body. Second return
// signals whether a usable schema exists.
func requestBodyProps(op *openapi3.Operation) (map[string]struct{}, bool) {
	if op == nil || op.RequestBody == nil || op.RequestBody.Value == nil {
		return nil, false
	}
	for _, mt := range op.RequestBody.Value.Content {
		if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
			continue
		}
		return schemaPropertyNames(mt.Schema.Value), true
	}
	return nil, false
}

// schemaPropertyNames collects top-level property keys from an OpenAPI
// schema, resolving one level of allOf so discriminated schemas work.
func schemaPropertyNames(s *openapi3.Schema) map[string]struct{} {
	out := map[string]struct{}{}
	if s == nil {
		return out
	}
	for name := range s.Properties {
		out[name] = struct{}{}
	}
	for _, ref := range s.AllOf {
		if ref == nil || ref.Value == nil {
			continue
		}
		for name := range ref.Value.Properties {
			out[name] = struct{}{}
		}
	}
	return out
}

// boolSet converts a struct{} set into a bool map so sortedKeys can be
// reused. A dedicated helper keeps the two callers symmetric.
func boolSet(s map[string]struct{}) map[string]bool {
	out := make(map[string]bool, len(s))
	for k := range s {
		out[k] = true
	}
	return out
}

// opLabel returns a reader-friendly identifier for an operation:
// operationId when present, otherwise the first tag, otherwise "op".
func opLabel(op *openapi3.Operation) string {
	if op == nil {
		return "op"
	}
	if op.OperationID != "" {
		return op.OperationID
	}
	if len(op.Tags) > 0 {
		return op.Tags[0]
	}
	return "op"
}

