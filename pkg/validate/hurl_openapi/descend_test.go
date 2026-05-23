//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what descend — schema를 한 세그먼트만큼 내려가는 로직 검증 (array/property/allOf)

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestDescend(t *testing.T) {
	nameSchema := &openapi3.Schema{Type: &openapi3.Types{"string"}}
	itemSchema := &openapi3.Schema{Type: &openapi3.Types{"object"}}

	cases := []struct {
		name    string
		schema  *openapi3.Schema
		seg     string
		wantNil bool
		wantPtr *openapi3.Schema
	}{
		{
			name:    "nil_schema",
			schema:  nil,
			seg:     "name",
			wantNil: true,
		},
		{
			name: "property_found",
			schema: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"name": &openapi3.SchemaRef{Value: nameSchema},
				},
			},
			seg:     "name",
			wantPtr: nameSchema,
		},
		{
			name: "property_not_found",
			schema: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"name": &openapi3.SchemaRef{Value: nameSchema},
				},
			},
			seg:     "email",
			wantNil: true,
		},
		{
			name: "array_index_descends_to_items",
			schema: &openapi3.Schema{
				Items: &openapi3.SchemaRef{Value: itemSchema},
			},
			seg:     "[0]",
			wantPtr: itemSchema,
		},
		{
			name: "array_index_nil_items",
			schema: &openapi3.Schema{
				Items: nil,
			},
			seg:     "[0]",
			wantNil: true,
		},
		{
			name: "allof_property_found",
			schema: &openapi3.Schema{
				AllOf: openapi3.SchemaRefs{
					{Value: &openapi3.Schema{
						Properties: openapi3.Schemas{
							"id": &openapi3.SchemaRef{Value: nameSchema},
						},
					}},
				},
			},
			seg:     "id",
			wantPtr: nameSchema,
		},
		{
			name: "allof_property_not_found",
			schema: &openapi3.Schema{
				AllOf: openapi3.SchemaRefs{
					{Value: &openapi3.Schema{
						Properties: openapi3.Schemas{
							"id": &openapi3.SchemaRef{Value: nameSchema},
						},
					}},
				},
			},
			seg:     "email",
			wantNil: true,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSchemaPointerCase(t, descend(c.schema, c.seg), c.wantNil, c.wantPtr)
		})
	}
}
