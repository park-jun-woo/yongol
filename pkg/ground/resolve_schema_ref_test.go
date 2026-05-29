//ff:func feature=rule type=test control=sequence
//ff:what resolveSchemaRef — object properties 분기

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
