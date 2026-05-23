//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what descendAllOf — allOf 멤버에서 seg property 탐색 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestDescendAllOf(t *testing.T) {
	nameSchema := &openapi3.Schema{Type: &openapi3.Types{"string"}}

	cases := []struct {
		name    string
		allOf   openapi3.SchemaRefs
		seg     string
		wantNil bool
		wantPtr *openapi3.Schema
	}{
		{
			name:    "nil_allof",
			allOf:   nil,
			seg:     "name",
			wantNil: true,
		},
		{
			name: "property_in_first_member",
			allOf: openapi3.SchemaRefs{
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"name": &openapi3.SchemaRef{Value: nameSchema},
					},
				}},
			},
			seg:     "name",
			wantPtr: nameSchema,
		},
		{
			name: "property_in_second_member",
			allOf: openapi3.SchemaRefs{
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"id": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					},
				}},
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"name": &openapi3.SchemaRef{Value: nameSchema},
					},
				}},
			},
			seg:     "name",
			wantPtr: nameSchema,
		},
		{
			name: "property_not_found",
			allOf: openapi3.SchemaRefs{
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"id": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					},
				}},
			},
			seg:     "email",
			wantNil: true,
		},
		{
			name: "nil_ref_skipped",
			allOf: openapi3.SchemaRefs{
				nil,
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"name": &openapi3.SchemaRef{Value: nameSchema},
					},
				}},
			},
			seg:     "name",
			wantPtr: nameSchema,
		},
		{
			name: "nil_value_skipped",
			allOf: openapi3.SchemaRefs{
				{Value: nil},
				{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"name": &openapi3.SchemaRef{Value: nameSchema},
					},
				}},
			},
			seg:     "name",
			wantPtr: nameSchema,
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runSchemaPointerCase(t, descendAllOf(c.allOf, c.seg), c.wantNil, c.wantPtr)
		})
	}
}
