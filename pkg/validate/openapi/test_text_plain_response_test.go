//ff:func feature=validate type=test-helper control=sequence topic=response-body-required
//ff:what 테스트 헬퍼 — content: text/plain (JSON 이 아닌) ResponseRef 빌드

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// textPlainResponse builds a ResponseRef with content: text/plain (not
// application/json). Used by O-5 unit tests to verify that non-JSON content
// types are rejected.
func textPlainResponse(desc string) *openapi3.ResponseRef {
	d := desc
	schema := openapi3.NewSchema()
	schema.Type = &openapi3.Types{"string"}
	return &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Description: &d,
			Content: openapi3.Content{
				"text/plain": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{Value: schema},
				},
			},
		},
	}
}
