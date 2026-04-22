//ff:func feature=gen-gogin type=util control=iteration dimension=1
//ff:what collectFrom200Response — 단일 operation의 200 응답 스키마 프로퍼티에서 $ref 이름 수집

package ssac

import "github.com/getkin/kin-openapi/openapi3"

// collectFrom200Response records every $ref schema name referenced (directly
// or as an array items ref) by a 200-response body's top-level properties.
// Silently returns for operations without a JSON 200 response schema.
func collectFrom200Response(op *openapi3.Operation, out map[string]bool) {
	if op == nil {
		return
	}
	resp := op.Responses.Status(200)
	if resp == nil || resp.Value == nil || resp.Value.Content == nil {
		return
	}
	mt := resp.Value.Content.Get("application/json")
	if mt == nil || mt.Schema == nil || mt.Schema.Value == nil {
		return
	}
	for _, propRef := range mt.Schema.Value.Properties {
		if name := extractRefName(propRef); name != "" {
			out[name] = true
		}
	}
}
