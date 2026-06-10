//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestCollectItemFieldTypes — 비배열/items 없음/프로퍼티 없음 nil 분기와 타입 맵 반환 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectItemFieldTypes(t *testing.T) {
	// scalar yields nil
	scalar := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	if collectItemFieldTypes(scalar) != nil {
		t.Error("scalar should yield nil")
	}
	// array with nil items yields nil
	noItems := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"array"}}}
	if collectItemFieldTypes(noItems) != nil {
		t.Error("array without items should yield nil")
	}
	// array whose items have no typed properties yields nil
	empty := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type:  &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"object"}}},
	}}
	if collectItemFieldTypes(empty) != nil {
		t.Error("items without typed properties should yield nil")
	}
	// typed properties are mapped name → type; nil-value and empty-Types
	// properties are skipped
	arr := &openapi3.SchemaRef{Value: &openapi3.Schema{
		Type: &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"object"},
			Properties: openapi3.Schemas{
				"id":       &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
				"name":     &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				"nilval":   {},
				"untyped":  &openapi3.SchemaRef{Value: &openapi3.Schema{}},
				"emptytyp": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{}}},
			},
		}},
	}}
	got := collectItemFieldTypes(arr)
	if got["id"] != "integer" || got["name"] != "string" {
		t.Errorf("collectItemFieldTypes = %v", got)
	}
	if len(got) != 2 {
		t.Errorf("nil/untyped/empty-typed properties must be skipped: %v", got)
	}
}
