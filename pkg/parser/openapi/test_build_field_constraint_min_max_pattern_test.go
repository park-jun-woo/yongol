//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestBuildFieldConstraintMinimumMaximum — Minimum, Maximum 필드 추출 검증

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
