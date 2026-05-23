//ff:func feature=validate type=test control=iteration dimension=1 topic=hurl-openapi
//ff:what jsonPathReachable — JSONPath가 OpenAPI schema에서 도달 가능한지 판정 검증

package hurl_openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestJsonPathReachable(t *testing.T) {
	// Build a schema: { user: { id: string, name: string }, items: [{ title: string }] }
	schema := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"user": &openapi3.SchemaRef{Value: &openapi3.Schema{
				Properties: openapi3.Schemas{
					"id":   &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
					"name": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				},
			}},
			"items": &openapi3.SchemaRef{Value: &openapi3.Schema{
				Type: &openapi3.Types{"array"},
				Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Properties: openapi3.Schemas{
						"title": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
					},
				}},
			}},
		},
	}

	cases := []struct {
		name string
		path string
		want bool
	}{
		{name: "nil_schema", path: "$.user.id", want: false},
		{name: "empty_path", path: "", want: false},
		{name: "direct_property", path: "$.user", want: true},
		{name: "nested_property", path: "$.user.id", want: true},
		{name: "nested_property2", path: "$.user.name", want: true},
		{name: "missing_property", path: "$.user.email", want: false},
		{name: "array_index_then_property", path: "$.items[0].title", want: true},
		{name: "array_index_missing_prop", path: "$.items[0].missing", want: false},
		{name: "wildcard_treated_reachable", path: "$..email", want: true},
		{name: "star_wildcard_treated_reachable", path: "$[*]", want: true},
		{name: "filter_treated_reachable", path: "$[?(@.active)]", want: true},
		{name: "top_level_missing", path: "$.nonexistent", want: false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s := schema
			if c.name == "nil_schema" {
				s = nil
			}
			got := jsonPathReachable(c.path, s)
			if got != c.want {
				t.Errorf("jsonPathReachable(%q) = %v, want %v", c.path, got, c.want)
			}
		})
	}
}
