//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-openapi
//ff:what withRequestBody — 테스트용: 주어진 properties 로 JSON requestBody 를 attach 한 operation

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

// withRequestBody attaches a JSON request-body schema with the given
// properties to a synthetic operation. Keeps fixtures inline.
func withRequestBody(opID string, props map[string]*openapi3.Schema) *openapi3.Operation {
	schema := &openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: openapi3.Schemas{},
	}
	for k, v := range props {
		schema.Properties[k] = &openapi3.SchemaRef{Value: v}
	}
	return &openapi3.Operation{
		OperationID: opID,
		Responses:   openapi3.NewResponses(),
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.Content{
					"application/json": &openapi3.MediaType{
						Schema: &openapi3.SchemaRef{Value: schema},
					},
				},
			},
		},
	}
}
