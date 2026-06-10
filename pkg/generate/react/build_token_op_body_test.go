//ff:func feature=gen-react type=test-helper control=iteration dimension=1
//ff:what 테스트 헬퍼 — JSON requestBody 스키마(string 프로퍼티 목록) 생성

package react

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// buildTokenOpBody builds a JSON requestBody whose object schema declares
// bodyFields (all string-typed). Split out of buildTokenOp.
func buildTokenOpBody(bodyFields []string) *openapi3.RequestBodyRef {
	bodyProps := openapi3.Schemas{}
	for _, f := range bodyFields {
		bodyProps[f] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	return &openapi3.RequestBodyRef{
		Value: &openapi3.RequestBody{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{Type: &openapi3.Types{"object"}, Properties: bodyProps},
					},
				},
			},
		},
	}
}
