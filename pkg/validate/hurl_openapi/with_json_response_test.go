//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-openapi
//ff:what withJSONResponse — 테스트용: status 에 JSON response body schema 를 attach 한 operation

package hurl_openapi

import "github.com/getkin/kin-openapi/openapi3"

// withJSONResponse attaches a JSON response-body schema to a synthetic
// operation. Keeps fixtures inline.
func withJSONResponse(opID, status string, props map[string]*openapi3.Schema) *openapi3.Operation {
	schema := &openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: openapi3.Schemas{},
	}
	for k, v := range props {
		schema.Properties[k] = &openapi3.SchemaRef{Value: v}
	}
	resp := &openapi3.Response{
		Content: openapi3.Content{
			"application/json": &openapi3.MediaType{
				Schema: &openapi3.SchemaRef{Value: schema},
			},
		},
	}
	op := &openapi3.Operation{
		OperationID: opID,
		Responses:   openapi3.NewResponses(),
	}
	op.Responses.Set(status, &openapi3.ResponseRef{Value: resp})
	return op
}
