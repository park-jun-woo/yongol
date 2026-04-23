//ff:func feature=openapi-parse type=test control=sequence
//ff:what ExtractRequestConstraints — nil doc 입력에서 빈 non-nil map 반환

package openapi

import "testing"

func TestExtractRequestConstraints_NilDoc(t *testing.T) {
	cs := ExtractRequestConstraints(nil, nil)
	if cs == nil {
		t.Fatal("nil doc should return empty non-nil map")
	}
	if len(cs) != 0 {
		t.Errorf("got %d entries, want 0", len(cs))
	}
}
