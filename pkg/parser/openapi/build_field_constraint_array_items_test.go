//ff:func feature=openapi-parse type=test control=sequence
//ff:what buildFieldConstraint가 array 타입의 items.type을 ItemType에 추출하는지 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildFieldConstraint_ArrayItemsString(t *testing.T) {
	prop := &openapi3.Schema{
		Type: &openapi3.Types{"array"},
		Items: &openapi3.SchemaRef{
			Value: &openapi3.Schema{
				Type: &openapi3.Types{"string"},
			},
		},
	}
	fc := buildFieldConstraint(prop, true)
	if fc.Type != "array" {
		t.Errorf("Type = %q, want \"array\"", fc.Type)
	}
	if fc.ItemType != "string" {
		t.Errorf("ItemType = %q, want \"string\"", fc.ItemType)
	}
}
