//ff:func feature=rule type=test control=sequence
//ff:what resolveSchemaRef — array items 분기: items.properties 를 반환

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

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
