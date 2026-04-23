//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectFrom200Response — 성공 2xx 응답의 직접/프로퍼티 $ref 이름 수집

package ssac

import (
	"github.com/getkin/kin-openapi/openapi3"

	yopenapi "github.com/park-jun-woo/yongol/pkg/parser/openapi"
)

// collectFrom200Response records every $ref schema name referenced by the
// operation's success response — either as a direct top-level $ref (the
// whole body is a single schema reference) or through its top-level
// properties' refs. Silently returns for operations without a JSON
// success response schema. The function name is historical; the selected
// status is DeriveSuccessStatus so POST handlers read 201, DELETE handlers
// read 204, etc. (BUG-004 ripple into converter collection).
//
// Direct $ref body collection was added for BUG-003: a plain
// `@response <var>` routes the sqlc row through convert<Model>, and
// convert<Model> must have been emitted regardless of whether the 2xx
// body schema has top-level properties.
func collectFrom200Response(op *openapi3.Operation, method string, out map[string]bool) {
	if op == nil {
		return
	}
	status := yopenapi.DeriveSuccessStatus(op, method)
	if status == 0 {
		return
	}
	resp := op.Responses.Status(status)
	if resp == nil || resp.Value == nil || resp.Value.Content == nil {
		return
	}
	mt := resp.Value.Content.Get("application/json")
	if mt == nil || mt.Schema == nil {
		return
	}
	// Direct body ref: `schema: { $ref: '#/components/schemas/Workflow' }`.
	if name := extractRefName(mt.Schema); name != "" {
		out[name] = true
	}
	if mt.Schema.Value == nil {
		return
	}
	for _, propRef := range mt.Schema.Value.Properties {
		if name := extractRefName(propRef); name != "" {
			out[name] = true
		}
	}
}
