//ff:func feature=openapi-parse type=test control=sequence
//ff:what collectItemFields/extractArrayItemFields/extractBodyConstraints/extractResponseFields/collect*ForOp/fill*Lines 직접 단위 검증
package openapi

import (
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
)

func TestCollectItemFields(t *testing.T) {
	got := collectItemFields(arraySchemaRef("id", "name"))
	if !got["id"] || !got["name"] {
		t.Errorf("collectItemFields = %v", got)
	}
	// non-array yields nil
	scalar := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"string"}}}
	if collectItemFields(scalar) != nil {
		t.Errorf("scalar should yield nil")
	}
	// array with nil items yields nil
	noItems := &openapi3.SchemaRef{Value: &openapi3.Schema{Type: &openapi3.Types{"array"}}}
	if collectItemFields(noItems) != nil {
		t.Errorf("array without items should yield nil")
	}
}
