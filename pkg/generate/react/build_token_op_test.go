//ff:func feature=gen-react type=test-helper control=iteration dimension=1
//ff:what 테스트 헬퍼 — 200 JSON 응답 필드와 requestBody 필드를 가진 operation 생성

package react

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildTokenOp builds an operation with a 200 application/json object
// response declaring respFields and, when bodyFields is non-empty, a JSON
// requestBody declaring bodyFields (all string-typed). Used by the Phase004
// refresh-plan and api-client tests.
func buildTokenOp(id string, respFields, bodyFields []string) *openapi3.Operation {
	respProps := openapi3.Schemas{}
	for _, f := range respFields {
		respProps[f] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{Type: &openapi3.Types{"object"}, Properties: respProps},
					},
				},
			},
		},
	})
	op := &openapi3.Operation{OperationID: id, Responses: responses}
	if len(bodyFields) > 0 {
		op.RequestBody = buildTokenOpBody(bodyFields)
	}
	return op
}
