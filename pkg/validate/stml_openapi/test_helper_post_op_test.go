//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what postOp — 테스트용 POST operation PathItem 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// postOp creates a PathItem with a POST operation.
func postOp(opID string, reqProps map[string]*openapi3.SchemaRef) *openapi3.PathItem {
	op := &openapi3.Operation{
		OperationID: opID,
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
					Type:       &openapi3.Types{"object"},
					Properties: reqProps,
				}),
			},
		},
		Responses: openapi3.NewResponses(),
	}
	return &openapi3.PathItem{Post: op}
}
