//ff:func feature=rule type=test control=sequence dimension=1
//ff:what resolveSchemaRef — 내부 properties, array items.properties, empty 대응

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestResolveSchemaRef_ObjectProperties returns top-level field names.
func TestResolveSchemaRef_ObjectProperties(t *testing.T) {
	prop := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"a": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"b": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	got := resolveSchemaRef(prop)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2: %v", len(got), got)
	}
}

// TestResolveSchemaRef_ArrayItems returns inner items' property names.
func TestResolveSchemaRef_ArrayItems(t *testing.T) {
	prop := &openapi3.Schema{
		Type: &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
			Type: &openapi3.Types{"object"},
			Properties: openapi3.Schemas{
				"id":   {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				"name": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			},
		}},
	}
	got := resolveSchemaRef(prop)
	if len(got) != 2 {
		t.Fatalf("array items len = %d, want 2: %v", len(got), got)
	}
}

// TestResolveSchemaRef_EmptyReturnsNil — no properties + not array = nil.
func TestResolveSchemaRef_EmptyReturnsNil(t *testing.T) {
	prop := &openapi3.Schema{Type: &openapi3.Types{"string"}}
	if got := resolveSchemaRef(prop); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}
