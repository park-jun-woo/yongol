//ff:func feature=openapi-parse type=test control=sequence
//ff:what buildFieldConstraint가 object 타입의 additionalProperties.type(integer)을 MapValueType에 추출하는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildFieldConstraint_MapValueInteger(t *testing.T) {
	prop := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		AdditionalProperties: openapi3.AdditionalProperties{
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}},
			},
		},
	}
	fc := buildFieldConstraint(prop, true)
	if fc.MapValueType != "integer" {
		t.Errorf("MapValueType = %q, want \"integer\"", fc.MapValueType)
	}
}
