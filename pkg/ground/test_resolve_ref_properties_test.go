//ff:func feature=rule type=test control=sequence dimension=1
//ff:what resolveRefProperties — 중첩 object/array의 내부 필드들을 평탄화해 반환

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestResolveRefProperties_ObjectAndArray combines a nested object property
// and an array-of-objects property to cover both resolveSchemaRef branches.
func TestResolveRefProperties_ObjectAndArray(t *testing.T) {
	schema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"item": {Value: &openapi3.Schema{
				Type: &openapi3.Types{"object"},
				Properties: openapi3.Schemas{
					"id":   {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
					"name": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
				},
			}},
			"items": {Value: &openapi3.Schema{
				Type: &openapi3.Types{"array"},
				Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Type: &openapi3.Types{"object"},
					Properties: openapi3.Schemas{
						"label": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
					},
				}},
			}},
			"count": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		},
	}
	got := resolveRefProperties(schema)

	seen := map[string]bool{}
	for _, f := range got {
		seen[f] = true
	}
	// Expected 3 nested fields flattened: id, name, label.
	// "count" has no nested properties, so contributes nothing.
	for _, f := range []string{"id", "name", "label"} {
		if !seen[f] {
			t.Errorf("resolved list missing %q: %v", f, got)
		}
	}
}
