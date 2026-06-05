//ff:func feature=openapi-parse type=test control=sequence
//ff:what buildFieldConstraint가 additionalProperties:true(Has) object를 자유형 마커 "*"로 추출하는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildFieldConstraint_FreeFormObjectHas(t *testing.T) {
	has := true
	prop := &openapi3.Schema{
		Type:                 &openapi3.Types{"object"},
		AdditionalProperties: openapi3.AdditionalProperties{Has: &has},
	}
	fc := buildFieldConstraint(prop, false)
	if fc.MapValueType != "*" {
		t.Errorf("MapValueType = %q, want \"*\" (free-form marker)", fc.MapValueType)
	}
}
