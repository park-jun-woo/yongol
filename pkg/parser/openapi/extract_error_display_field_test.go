//ff:func feature=openapi-parse type=test control=iteration dimension=1
//ff:what ExtractErrorDisplayField — error/message/없음/비-string/nil 스키마별 표시 필드 도출 검증
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractErrorDisplayField(t *testing.T) {
	str := func() *openapi3.SchemaRef {
		return &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	}
	doc := func(props openapi3.Schemas) *openapi3.T {
		return &openapi3.T{
			Components: &openapi3.Components{
				Schemas: openapi3.Schemas{
					"ErrorResponse": {Value: &openapi3.Schema{
						Type:       &openapi3.Types{"object"},
						Properties: props,
					}},
				},
			},
		}
	}

	cases := []struct {
		name  string
		input *openapi3.T
		want  string
	}{
		{"error+code", doc(openapi3.Schemas{"error": str(), "code": str()}), "error"},
		{"message only", doc(openapi3.Schemas{"message": str()}), "message"},
		{"neither", doc(openapi3.Schemas{"detail": str()}), "error"},
		{"non-string error", doc(openapi3.Schemas{"error": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}}}), "error"},
		{"nil doc", nil, "error"},
	}
	for _, c := range cases {
		if got := ExtractErrorDisplayField(c.input); got != c.want {
			t.Errorf("%s: got %q, want %q", c.name, got, c.want)
		}
	}
}
