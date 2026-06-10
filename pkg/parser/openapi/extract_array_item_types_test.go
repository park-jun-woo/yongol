//ff:func feature=openapi-parse type=test control=sequence
//ff:what extractArrayItemTypes — 배열 프로퍼티만 항목 타입 맵으로 수집됨을 검증

package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractArrayItemTypes(t *testing.T) {
	schema := &openapi3.Schema{
		Type: &openapi3.Types{"object"},
		Properties: openapi3.Schemas{
			"photos": {Value: &openapi3.Schema{
				Type: &openapi3.Types{"array"},
				Items: &openapi3.SchemaRef{Value: &openapi3.Schema{
					Type: &openapi3.Types{"object"},
					Properties: openapi3.Schemas{
						"id": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
					},
				}},
			}},
			"total": {Value: &openapi3.Schema{Type: &openapi3.Types{"integer"}}},
		},
	}
	got := extractArrayItemTypes(schema)
	if len(got) != 1 {
		t.Fatalf("expected 1 array field, got %v", got)
	}
	if got["photos"]["id"] != "integer" {
		t.Errorf("photos.id = %q, want integer", got["photos"]["id"])
	}
}
