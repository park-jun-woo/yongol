//ff:func feature=manifest type=test control=sequence
//ff:what buildFieldConstraint가 Type.Slice() 빈 배열에서 panic 없이 빈 Type을 반환하는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildFieldConstraint_EmptyTypeSlice(t *testing.T) {
	// Types nil → Slice() returns nil → len == 0.
	prop := &openapi3.Schema{}
	fc := buildFieldConstraint(prop, true)
	if fc.Type != "" {
		t.Errorf("Type = %q, want \"\"", fc.Type)
	}
	if !fc.Required {
		t.Errorf("Required = false, want true")
	}

	// Explicitly empty Types.
	empty := &openapi3.Types{}
	prop2 := &openapi3.Schema{Type: empty}
	fc2 := buildFieldConstraint(prop2, false)
	if fc2.Type != "" {
		t.Errorf("Type(empty) = %q, want \"\"", fc2.Type)
	}
}
