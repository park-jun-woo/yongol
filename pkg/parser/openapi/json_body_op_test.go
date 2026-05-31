//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func jsonBodyOp(opID string, props openapi3.Schemas) *openapi3.Operation {
	schema := &openapi3.Schema{Properties: props}
	rb := openapi3.NewRequestBody().WithJSONSchema(schema)
	op := openapi3.NewOperation()
	op.OperationID = opID
	op.RequestBody = &openapi3.RequestBodyRef{Value: rb}
	return op
}
