//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestExtractArrayItemFields(t *testing.T) {
	schema := &openapi3.Schema{Properties: openapi3.Schemas{
		"items":  arraySchemaRef("id", "title"),
		"single": {Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}},
	}}
	got := extractArrayItemFields(schema)
	if len(got) != 1 {
		t.Fatalf("expected 1 array field, got %v", got)
	}
	if !got["items"]["id"] || !got["items"]["title"] {
		t.Errorf("items fields = %v", got["items"])
	}
}
