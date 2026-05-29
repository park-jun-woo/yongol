//ff:func feature=rule type=test control=sequence dimension=1
//ff:what populateOpenAPIParams — operationId 있는 operation의 param/request를 등록

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestPopulateOpenAPIParams_WithOpID covers happy path: params + requestBody
// are registered keyed by operationId.
func TestPopulateOpenAPIParams_WithOpID(t *testing.T) {
	param := &openapi3.ParameterRef{Value: &openapi3.Parameter{
		Name: "id",
		In:   "path",
	}}
	reqSchema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"title": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"body":  {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	op := &openapi3.Operation{
		OperationID: "UpdateItem",
		Parameters:  openapi3.Parameters{param},
		RequestBody: &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
			WithContent(openapi3.NewContentWithJSONSchema(reqSchema))},
	}
	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items/{id}", &openapi3.PathItem{Put: op}),
	)}
	fs := newMinimalFullstack(withOpenAPIDoc(doc))
	g := newGround()

	populateOpenAPIParams(g, fs)

	params := g.Lookup["OpenAPI.param.UpdateItem"]
	if !params["id"] {
		t.Errorf("OpenAPI.param.UpdateItem missing 'id': %v", params)
	}

	req := g.Lookup["OpenAPI.request.UpdateItem"]
	if !req["title"] || !req["body"] {
		t.Errorf("OpenAPI.request.UpdateItem missing title/body: %v", req)
	}
}
