//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestBuildFieldConstraintPattern — Pattern 필드 추출 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

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
