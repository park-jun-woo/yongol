//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what getOp — 테스트용 GET operation PathItem 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// getOp creates a PathItem with a GET operation.
func getOp(opID string, params []*openapi3.ParameterRef, respProps map[string]*openapi3.SchemaRef) *openapi3.PathItem {
	op := &openapi3.Operation{OperationID: opID, Parameters: params}
	if respProps != nil {
		op.Responses = openapi3.NewResponses(
			openapi3.WithStatus(200, &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
						Type:       &openapi3.Types{"object"},
						Properties: respProps,
					}),
				},
			}),
		)
	}
	return &openapi3.PathItem{Get: op}
}
