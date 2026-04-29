//ff:func feature=validate type=test-helper control=sequence topic=response-body-required
//ff:what 테스트 헬퍼 — content: application/json + 인라인 schema 보유 ResponseRef 빌드

package openapi

import "github.com/getkin/kin-openapi/openapi3"

// jsonResponseWithSchema builds a ResponseRef with content: application/json
// + inline object schema. Used by O-5 unit tests.
func jsonResponseWithSchema(desc string) *openapi3.ResponseRef {
	d := desc
	schema := openapi3.NewSchema()
	schema.Type = &openapi3.Types{"object"}
	return &openapi3.ResponseRef{
		Value: &openapi3.Response{
			Description: &d,
			Content: openapi3.Content{
				"application/json": &openapi3.MediaType{
					Schema: &openapi3.SchemaRef{Value: schema},
				},
			},
		},
	}
}
