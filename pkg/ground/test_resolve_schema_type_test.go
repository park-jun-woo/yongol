//ff:func feature=rule type=test control=sequence dimension=1
//ff:what resolveSchemaType — $ref 정규화 + primitive fallback

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestResolveSchemaType_Ref returns the last segment of a $ref path.
func TestResolveSchemaType_Ref(t *testing.T) {
	ref := &openapi3.SchemaRef{Ref: "#/components/schemas/Workflow"}
	if got := resolveSchemaType(ref); got != "Workflow" {
		t.Errorf("got %q, want Workflow", got)
	}
}

// TestResolveSchemaType_Primitive falls through to resolvePrimitiveType.
func TestResolveSchemaType_Primitive(t *testing.T) {
	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}}
	if got := resolveSchemaType(ref); got != "int" {
		t.Errorf("got %q, want int", got)
	}
}

// TestResolveSchemaType_Nil returns "".
func TestResolveSchemaType_Nil(t *testing.T) {
	if got := resolveSchemaType(nil); got != "" {
		t.Errorf("got %q, want empty", got)
	}
}
