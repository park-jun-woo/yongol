//ff:func feature=rule type=test control=iteration dimension=1
//ff:what extractRequestFields — multipart/form-data content type 필드 추출 검증

package ground

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractRequestFields_MultipartFormData(t *testing.T) {
	schema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"file":     {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}, Format: "binary"}},
			"category": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
		},
	}
	body := openapi3.NewRequestBody().WithContent(openapi3.Content{
		"multipart/form-data": openapi3.NewMediaType().WithSchema(schema),
	})

	got := extractRequestFields(body)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	for _, f := range []string{"file", "category"} {
		if !got[f] {
			t.Errorf("missing %q: %v", f, got)
		}
	}
}
