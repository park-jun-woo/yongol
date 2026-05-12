//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestExtractResponseArrayItemFields_NilDoc — nil doc 검증

package openapi

import "testing"

func TestExtractResponseArrayItemFields_NilDoc(t *testing.T) {
	result := ExtractResponseArrayItemFields(nil)
	if len(result) != 0 {
		t.Errorf("expected empty map for nil doc, got %v", result)
	}
}
