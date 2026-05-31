//ff:func feature=generate type=test control=sequence
//ff:what TestGenerateBatch_ZeroCov — pkg/generate 소형 순수/IO 헬퍼 분기 커버
package generate

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestBuildDefaultFieldConstraints_ZeroCov(t *testing.T) {
	// nil schema
	if buildDefaultFieldConstraints(nil) != nil {
		t.Error("nil schema should return nil")
	}
	// empty properties
	if buildDefaultFieldConstraints(&openapi3.Schema{}) != nil {
		t.Error("empty props should return nil")
	}
	// real schema with required + a nil-value prop skipped
	schema := &openapi3.Schema{
		Required: []string{"email"},
		Properties: openapi3.Schemas{
			"email": &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "email"}},
			"bad":   &openapi3.SchemaRef{}, // nil Value → skipped
		},
	}
	got := buildDefaultFieldConstraints(schema)
	if got == nil || len(got) != 1 {
		t.Fatalf("expected 1 field, got %v", got)
	}
	if !got["email"].Required || got["email"].Type != "string" || got["email"].Format != "email" {
		t.Errorf("email constraint = %+v", got["email"])
	}
}
