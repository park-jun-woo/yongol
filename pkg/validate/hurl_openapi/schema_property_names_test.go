//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what schemaPropertyNames — schema top-level + allOf members의 property 이름 집합 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSchemaPropertyNames(t *testing.T) {
	cases := []struct {
		name string
		s    *openapi3.Schema
		want map[string]struct{}
	}{
		{name: "nil_schema", s: nil, want: map[string]struct{}{}},
		{
			name: "direct_properties",
			s: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id":   &openapi3.SchemaRef{Value: &openapi3.Schema{}},
					"name": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				},
			},
			want: map[string]struct{}{"id": {}, "name": {}},
		},
		{
			name: "allof_properties_merged",
			s: &openapi3.Schema{
				AllOf: openapi3.SchemaRefs{
					{Value: &openapi3.Schema{
						Properties: openapi3.Schemas{
							"email": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
						},
					}},
				},
			},
			want: map[string]struct{}{"email": {}},
		},
		{
			name: "mixed_properties_and_allof",
			s: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				},
				AllOf: openapi3.SchemaRefs{
					{Value: &openapi3.Schema{
						Properties: openapi3.Schemas{
							"email": &openapi3.SchemaRef{Value: &openapi3.Schema{}},
						},
					}},
				},
			},
			want: map[string]struct{}{"id": {}, "email": {}},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			runStructSetCase(t, schemaPropertyNames(c.s), c.want)
		})
	}
}
