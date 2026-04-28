//ff:func feature=validate type=util control=sequence topic=hurl-openapi
//ff:what opLabel — operation 의 reader-friendly 식별자 (operationId → tag → "op")

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

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
