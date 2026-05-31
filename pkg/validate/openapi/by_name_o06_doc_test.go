//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what TestByName_ZeroCov — O-6 스키마 워커들을 이름으로 직접 호출해 커버리지 귀속
package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

func byNameO06Doc() *openapi3.T {
	itemSchema := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:       &openapi3.Types{"object"},
		Properties: openapi3.Schemas{"id": {Value: openapi3.NewSchema()}},
		Required:   []string{"phantom"},
	}}
	arraySchema := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: itemSchema,
	}}
	objSchema := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"name":  {Value: openapi3.NewSchema()},
			"items": arraySchema,
		},
		Required: []string{"name", "missing"},
	}}

	reqBody := &openapi3.RequestBodyRef{Value: openapi3.NewRequestBody().
		WithJSONSchemaRef(objSchema)}

	resp := openapi3.NewResponse().WithJSONSchemaRef(arraySchema)
	resps := openapi3.NewResponses()
	resps.Set("200", &openapi3.ResponseRef{Value: resp})

	op := &openapi3.Operation{OperationID: "CreateItem", RequestBody: reqBody, Responses: resps}
	pi := &openapi3.PathItem{Post: op}

	return &openapi3.T{
		Components: &openapi3.Components{Schemas: openapi3.Schemas{"Workflow": objSchema}},
		Paths:      openapi3.NewPaths(openapi3.WithPath("/items", pi)),
	}
}
