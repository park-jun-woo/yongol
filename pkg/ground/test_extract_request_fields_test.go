//ff:func feature=rule type=test control=sequence dimension=1
//ff:what extractRequestFields — requestBody JSON 스키마 → field 이름 셋

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRequestFields_JSONSchema(t *testing.T) {
	schema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"title":    {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"body":     {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
			"priority": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		},
	}
	body := openapi3.NewRequestBody().WithContent(openapi3.NewContentWithJSONSchema(schema))

	got := extractRequestFields(body)
	if len(got) != 3 {
		t.Fatalf("len = %d, want 3", len(got))
	}
	for _, f := range []string{"title", "body", "priority"} {
		if !got[f] {
			t.Errorf("missing %q: %v", f, got)
		}
	}
}

func TestExtractRequestFields_NilContent(t *testing.T) {
	body := &openapi3.RequestBody{}
	if got := extractRequestFields(body); got != nil {
		t.Errorf("expected nil for empty body, got %v", got)
	}
}
