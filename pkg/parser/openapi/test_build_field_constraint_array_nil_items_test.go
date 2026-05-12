//ff:func feature=openapi-parse type=test control=sequence
//ff:what buildFieldConstraint가 array 타입에서 Items nil일 때 ItemType 빈 문자열 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildFieldConstraint_ArrayNilItems(t *testing.T) {
	prop := &openapi3.Schema{
		Type: &openapi3.Types{"array"},
	}
	fc := buildFieldConstraint(prop, true)
	if fc.Type != "array" {
		t.Errorf("Type = %q, want \"array\"", fc.Type)
	}
	if fc.ItemType != "" {
		t.Errorf("ItemType = %q, want \"\"", fc.ItemType)
	}
}
