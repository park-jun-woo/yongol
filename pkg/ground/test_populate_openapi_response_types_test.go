//ff:func feature=rule type=test control=sequence
//ff:what populateOpenAPIResponseTypes — 첫 2xx response의 각 field 타입 등록

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestPopulateOpenAPIResponseTypes_FirstTwoXX verifies Types[OpenAPI.response.
// <opID>.<field>] is populated from the first 2xx schema, mapping primitive
// types via resolveSchemaType.
func TestPopulateOpenAPIResponseTypes_FirstTwoXX(t *testing.T) {
	op := &openapi3.Operation{OperationID: "GetItem"}
	respSchema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"id":   {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}, Format: "int64"}},
			"name": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	resp := openapi3.NewResponse().WithContent(openapi3.NewContentWithJSONSchema(respSchema))
	op.Responses = openapi3.NewResponses()
	op.Responses.Set("200", &openapi3.ResponseRef{Value: resp})

	doc := &openapi3.T{Paths: openapi3.NewPaths(
		openapi3.WithPath("/items", &openapi3.PathItem{Get: op}),
	)}
	fs := newMinimalFullstack(withOpenAPIDoc(doc))
	g := newGround()

	populateOpenAPIResponseTypes(g, fs)

	if g.Types["OpenAPI.response.GetItem.id"] != "int64" {
		t.Errorf("id type = %q, want int64", g.Types["OpenAPI.response.GetItem.id"])
	}
	if g.Types["OpenAPI.response.GetItem.name"] != "string" {
		t.Errorf("name type = %q, want string", g.Types["OpenAPI.response.GetItem.name"])
	}
}
