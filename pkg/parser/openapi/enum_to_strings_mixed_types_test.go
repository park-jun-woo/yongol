//ff:func feature=openapi-parse type=test control=iteration dimension=1
//ff:what enumToStrings — string/int/bool 혼합 Enum 변환

package openapi

import "testing"

func TestEnumToStrings_MixedTypes(t *testing.T) {
	got := enumToStrings([]any{"a", 1, true})
	want := []string{"a", "1", "true"}
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d", len(got), len(want))
	}
	for i, s := range want {
		if got[i] != s {
			t.Errorf("[%d] = %q, want %q", i, got[i], s)
		}
	}
}
