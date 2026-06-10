//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what postOpWithResp — 200 응답 스키마를 가진 테스트용 POST operation PathItem 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// postOpWithResp creates a PathItem with a POST operation whose 200
// response declares the given top-level properties.
func postOpWithResp(opID string, respProps map[string]*openapi3.SchemaRef) *openapi3.PathItem {
	op := &openapi3.Operation{
		OperationID: opID,
		Responses: openapi3.NewResponses(
			openapi3.WithStatus(200, &openapi3.ResponseRef{
				Value: &openapi3.Response{
					Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
						Type:       &openapi3.Types{"object"},
						Properties: respProps,
					}),
				},
			}),
		),
	}
	return &openapi3.PathItem{Post: op}
}
