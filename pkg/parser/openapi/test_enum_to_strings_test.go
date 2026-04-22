//ff:func feature=manifest type=test control=sequence
//ff:what enumToStrings / buildFieldConstraint Enum 경로 회귀

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestEnumToStrings_MixedTypes(t *testing.T) {
	got := enumToStrings([]any{"a", 1, true})
	want := []string{"a", "1", "true"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i, s := range want {
		if got[i] != s {
			t.Errorf("[%d] = %q, want %q", i, got[i], s)
		}
	}
}

func TestEnumToStrings_Empty(t *testing.T) {
	if got := enumToStrings(nil); got != nil {
		t.Errorf("got %v, want nil", got)
	}
}

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
