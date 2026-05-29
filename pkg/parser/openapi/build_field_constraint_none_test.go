//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestBuildFieldConstraintNoMinMaxPattern — constraint 없는 스키마 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildFieldConstraintNoMinMaxPattern(t *testing.T) {
	prop := &openapi3.Schema{
		Type: &openapi3.Types{"string"},
	}
	fc := buildFieldConstraint(prop, false)
	if fc.Minimum != nil {
		t.Errorf("expected nil Minimum, got %v", fc.Minimum)
	}
	if fc.Maximum != nil {
		t.Errorf("expected nil Maximum, got %v", fc.Maximum)
	}
	if fc.Pattern != "" {
		t.Errorf("expected empty Pattern, got %q", fc.Pattern)
	}
}
