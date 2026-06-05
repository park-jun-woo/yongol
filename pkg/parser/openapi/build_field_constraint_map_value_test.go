//ff:func feature=openapi-parse type=test control=sequence
//ff:what buildFieldConstraint가 object 타입의 additionalProperties.type(string)을 MapValueType에 추출하는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildFieldConstraint_MapValueString(t *testing.T) {
	prop := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		AdditionalProperties: openapi3.AdditionalProperties{
			Schema: &openapi3.SchemaRef{
				Value: &openapi3.Schema{Type: &openapi3.Types{"string"}},
			},
		},
	}
	fc := buildFieldConstraint(prop, true)
	if fc.Type != "object" {
		t.Errorf("Type = %q, want \"object\"", fc.Type)
	}
	if fc.MapValueType != "string" {
		t.Errorf("MapValueType = %q, want \"string\"", fc.MapValueType)
	}
}
