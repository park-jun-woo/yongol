//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=config-check
//ff:what 테스트 헬퍼 — 필드명 목록으로 200 application/json object 응답 생성

package openapi_manifest

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// build2xxObjectResponses builds a Responses with a single 200
// application/json object schema whose top-level properties are the given
// field names (all string-typed). Used by buildDocWithResponseFields.
func build2xxObjectResponses(fields []string) *openapi3.Responses {
	props := openapi3.Schemas{}
	for _, f := range fields {
		props[f] = &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	responses := openapi3.NewResponses()
	responses.Set("200", &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{
						Value: &openapi3.Schema{Type: &openapi3.Types{"object"}, Properties: props},
					},
				},
			},
		},
	})
	return responses
}
