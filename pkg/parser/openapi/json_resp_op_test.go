//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func jsonRespOp(opID string, props openapi3.Schemas) *openapi3.Operation {
	op := openapi3.NewOperation()
	op.OperationID = opID
	resp := openapi3.NewResponse().WithJSONSchema(&openapi3.Schema{Properties: props})
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{Value: resp})
	op.Responses = responses
	return op
}
