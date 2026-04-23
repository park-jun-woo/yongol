//ff:func feature=rule type=test control=sequence
//ff:what resolveSchemaType — primitive 분기: resolvePrimitiveType fallthrough

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestResolveSchemaType_Primitive falls through to resolvePrimitiveType.
func TestResolveSchemaType_Primitive(t *testing.T) {
	ref := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}}
	if got := resolveSchemaType(ref); got != "int" {
		t.Errorf("got %q, want int", got)
	}
}
