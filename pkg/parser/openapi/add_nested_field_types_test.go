//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestAddNestedFieldTypes — object/array 전개 + Type nil/array nil items 스킵 분기 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestAddNestedFieldTypes(t *testing.T) {
	// Type nil → no expansion, no panic.
	out := make(map[string]FieldTypeInfo)
	addNestedFieldTypes(out, "x", &openapi3.Schema{})
	if len(out) != 0 {
		t.Errorf("typeless schema must expand nothing, got %v", out)
	}

	// object → child props expanded.
	obj := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"name": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	addNestedFieldTypes(out, "user", obj)
	if out["user.name"].Type != "string" {
		t.Errorf("user.name: want string, got %+v", out["user.name"])
	}

	// array with items → item props expanded.
	arr := &openapi3.Schema{
		Type: &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"object"},
			Properties: openapi3.Schemas{
				"url": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			},
		}},
	}
	addNestedFieldTypes(out, "photos", arr)
	if out["photos.url"].Type != "string" {
		t.Errorf("photos.url: want string, got %+v", out["photos.url"])
	}

	// array without items → no expansion, no panic.
	before := len(out)
	addNestedFieldTypes(out, "empty", &openapi3.Schema{Type: &openapi3.Types{"array"}})
	if len(out) != before {
		t.Errorf("array without items must expand nothing")
	}
}
