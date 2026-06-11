//ff:func feature=validate type=test control=sequence topic=openapi-structural
//ff:what schemaPropIsString — string 속성 존재/타입 불일치/부재 판정 검증
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestSchemaPropIsString(t *testing.T) {
	schema := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"error": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"count": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			"nilv":  {},
		},
	}
	if !schemaPropIsString(schema, "error") {
		t.Errorf("error should be a string property")
	}
	if schemaPropIsString(schema, "count") {
		t.Errorf("count is integer, not a string property")
	}
	if schemaPropIsString(schema, "nilv") {
		t.Errorf("nil-value property is not a string property")
	}
	if schemaPropIsString(schema, "absent") {
		t.Errorf("absent property is not a string property")
	}
}
