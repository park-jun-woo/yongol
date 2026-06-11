//ff:func feature=validate type=test-helper control=sequence topic=stml-openapi
//ff:what editPath — GET-by-id + PUT/PATCH(requestBody)를 한 경로에 묶은 테스트용 PathItem 생성

package stml_openapi

import "github.com/getkin/kin-openapi/openapi3"

// editPath builds a PathItem carrying a GET (getResp 2xx props) plus a write
// operation under method ("PUT" or "PATCH") whose requestBody declares reqProps
// with the given required list. getOpID/writeOpID name the two operations.
func editPath(getOpID string, getResp map[string]*openapi3.SchemaRef, method, writeOpID string, reqProps map[string]*openapi3.SchemaRef, required []string) *openapi3.PathItem {
	get := &openapi3.Operation{
		OperationID: getOpID,
		Responses: openapi3.NewResponses(openapi3.WithStatus(200, &openapi3.ResponseRef{
			Value: &openapi3.Response{Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
				Type:       &openapi3.Types{"object"},
				Properties: getResp,
			})},
		})),
	}
	write := &openapi3.Operation{
		OperationID: writeOpID,
		RequestBody: &openapi3.RequestBodyRef{Value: &openapi3.RequestBody{
			Content: openapi3.NewContentWithJSONSchema(&openapi3.Schema{
				Type:       &openapi3.Types{"object"},
				Properties: reqProps,
				Required:   required,
			}),
		}},
		Responses: openapi3.NewResponses(),
	}
	item := &openapi3.PathItem{Get: get}
	if method == "PATCH" {
		item.Patch = write
	} else {
		item.Put = write
	}
	return item
}
