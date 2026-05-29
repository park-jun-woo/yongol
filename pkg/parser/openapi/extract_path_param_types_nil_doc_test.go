//ff:func feature=openapi-parse type=test control=sequence
//ff:what TestExtractPathParamTypes_NilDoc — nil doc에서 빈 맵을 반환하는지 검증

package openapi

import "testing"

func TestExtractPathParamTypes_NilDoc(t *testing.T) {
	result := ExtractPathParamTypes(nil)
	if len(result) != 0 {
		t.Errorf("len = %d, want 0", len(result))
	}
}
