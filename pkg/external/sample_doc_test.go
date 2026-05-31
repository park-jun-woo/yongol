//ff:func feature=external type=test control=sequence
//ff:what TestExtractMethods/buildMethodInfo/extractResponseTypes — 메서드·응답타입 추출 검증
package external

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func sampleDoc() *openapi3.T {
	getResp := openapi3.NewResponse().WithJSONSchema(&openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"id":   intSchema(),
			"name": strSchema(),
		},
	})
	getOp := &openapi3.Operation{
		OperationID: "get_item",
		Parameters: openapi3.Parameters{
			{Value: &openapi3.Parameter{Name: "item_id", In: "path", Schema: intSchema()}},
		},
		Responses: openapi3.NewResponses(openapi3.WithStatus(200, &openapi3.ResponseRef{Value: getResp})),
	}
	postOp := &openapi3.Operation{
		OperationID: "create_item",
		RequestBody: &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().WithJSONSchema(&openapi3.Schema{
			Type:       &openapi3.Types{"object"},
			Properties: openapi3.Schemas{"name": strSchema()},
		})},
		Responses: openapi3.NewResponses(), // no 200 => error-only
	}

	doc := &openapi3.T{Paths: openapi3.NewPaths()}
	doc.Paths.Set("/items/{item_id}", &openapi3.PathItem{Get: getOp})
	doc.Paths.Set("/items", &openapi3.PathItem{Post: postOp})
	return doc
}
