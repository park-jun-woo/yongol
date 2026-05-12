//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestExtractNoBodyOps_NilDoc — nil doc 검증

package openapi

import "testing"

func TestExtractNoBodyOps_NilDoc(t *testing.T) {
	result := ExtractNoBodyOps(nil)
	if len(result) != 0 {
		t.Errorf("expected empty map for nil doc, got %v", result)
	}
}
