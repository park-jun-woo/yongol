//ff:func feature=rule type=test control=sequence
//ff:what resolveSchemaRef — properties 없고 array 도 아닐 때 nil 반환

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

// TestResolveSchemaRef_EmptyReturnsNil — no properties + not array = nil.
func TestResolveSchemaRef_EmptyReturnsNil(t *testing.T) {
	prop := &openapi3.Schema{Type: &openapi3.Types{"string"}}
	if got := resolveSchemaRef(prop); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}
