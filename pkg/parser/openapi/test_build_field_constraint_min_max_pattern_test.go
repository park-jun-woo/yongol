//ff:func feature=openapi-parse type=test control=sequence
//ff:what buildFieldConstraint — Minimum, Maximum, Pattern 필드 추출을 검증
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildFieldConstraintMinimumMaximum(t *testing.T) {
	min := 1.0
	max := 100.0
	prop := &openapi3.Schema{
		Type: &openapi3.Types{"integer"},
		Min:  &min,
		Max:  &max,
	}
	fc := buildFieldConstraint(prop, true)
	if fc.Minimum == nil || *fc.Minimum != 1.0 {
		t.Errorf("expected Minimum=1, got %v", fc.Minimum)
	}
	if fc.Maximum == nil || *fc.Maximum != 100.0 {
		t.Errorf("expected Maximum=100, got %v", fc.Maximum)
	}
}

func TestBuildFieldConstraintPattern(t *testing.T) {
	prop := &openapi3.Schema{
		Type:    &openapi3.Types{"string"},
		Pattern: `^\d{3}$`,
	}
	fc := buildFieldConstraint(prop, false)
	if fc.Pattern != `^\d{3}$` {
		t.Errorf("expected Pattern=^\\d{3}$, got %q", fc.Pattern)
	}
}

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
