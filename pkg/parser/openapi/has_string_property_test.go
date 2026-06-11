//ff:func feature=openapi-parse type=test control=sequence
//ff:what hasStringProperty — string 속성 존재/타입 불일치/부재 판정 검증
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestHasStringProperty(t *testing.T) {
	schema := &openapi3.Schema{
		Properties: openapi3.Schemas{
			"error": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"count": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
			"nilv":  {},
		},
	}
	if !hasStringProperty(schema, "error") {
		t.Errorf("error should be a string property")
	}
	if hasStringProperty(schema, "count") {
		t.Errorf("count is integer, not a string property")
	}
	if hasStringProperty(schema, "nilv") {
		t.Errorf("nil-value property is not a string property")
	}
	if hasStringProperty(schema, "absent") {
		t.Errorf("absent property is not a string property")
	}
}
