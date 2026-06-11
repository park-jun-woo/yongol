//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestAddChildFieldTypes — 정상 하위 프로퍼티 기록 + nil ref/nil Value 스킵 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestAddChildFieldTypes(t *testing.T) {
	out := make(map[string]FieldTypeInfo)
	props := openapi3.Schemas{
		"good":     &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		"nilRef":   nil,
		"nilValue": &openapi3.SchemaRef{},
	}
	addChildFieldTypes(out, "parent", props)

	if out["parent.good"].Type != "string" {
		t.Errorf("parent.good: want string, got %+v", out["parent.good"])
	}
	if _, ok := out["parent.nilRef"]; ok {
		t.Errorf("nilRef must be skipped")
	}
	if _, ok := out["parent.nilValue"]; ok {
		t.Errorf("nilValue must be skipped")
	}
}
