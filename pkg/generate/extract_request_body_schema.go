//ff:func feature=generate type=util control=iteration dimension=1
//ff:what operation의 requestBody에서 스키마를 추출한다 (application/json 우선, 없으면 다른 content type 시도)

package generate

import "github.com/getkin/kin-openapi/openapi3"

// extractRequestBodySchema returns the resolved schema from the operation's
// requestBody. It checks application/json first and falls back to any other
// content type that carries a schema.
func extractRequestBodySchema(op *openapi3.Operation) *openapi3.Schema {
	if op.RequestBody == nil || op.RequestBody.Value == nil || op.RequestBody.Value.Content == nil {
		return nil
	}
	ct := op.RequestBody.Value.Content.Get("application/json")
	if ct != nil && ct.Schema != nil && ct.Schema.Value != nil {
		return ct.Schema.Value
	}
	// Fallback: try any content type that has a schema.
	for _, mt := range op.RequestBody.Value.Content {
		if mt.Schema != nil && mt.Schema.Value != nil {
			return mt.Schema.Value
		}
	}
	return nil
}
