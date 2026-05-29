//ff:func feature=rule type=test control=sequence
//ff:what resolveSchemaType — $ref 분기: 마지막 segment 를 반환

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
