//ff:func feature=openapi-parse type=test control=sequence
//ff:what buildFieldConstraint — Enum, MaxLength, MinLength, Format 동시 추출

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildFieldConstraint_EnumAndLengths(t *testing.T) {
	ml := uint64(64)
	prop := &openapi3.Schema{
		Type:      &openapi3.Types{"string"},
		Format:    "email",
		MaxLength: &ml,
		MinLength: 8,
		Enum:      []any{"draft", "published"},
	}
	fc := buildFieldConstraint(prop, false)
	if fc.Type != "string" || fc.Format != "email" {
		t.Errorf("Type/Format = %q/%q", fc.Type, fc.Format)
	}
	if fc.MaxLength == nil || *fc.MaxLength != 64 {
		t.Errorf("MaxLength = %v", fc.MaxLength)
	}
	if fc.MinLength == nil || *fc.MinLength != 8 {
		t.Errorf("MinLength = %v", fc.MinLength)
	}
	if len(fc.Enum) != 2 {
		t.Errorf("Enum = %v", fc.Enum)
	}
}
