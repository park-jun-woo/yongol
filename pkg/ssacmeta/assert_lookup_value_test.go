//ff:func feature=ssacmeta type=test-helper control=selection
//ff:what assertLookupValue — lookupPath 반환값을 want 타입(string/bool/map)에 맞춰 비교 헬퍼
package ssacmeta

import "testing"

// assertLookupValue compares a lookupPath result against want, dispatching on
// the want type (string/bool compared directly, map only checked non-nil).
func assertLookupValue(t *testing.T, path string, got, want any) {
	t.Helper()
	switch want := want.(type) {
	case string:
		if got != want {
			t.Errorf("lookupPath(%q) = %v, want %v", path, got, want)
		}
	case bool:
		if got != want {
			t.Errorf("lookupPath(%q) = %v, want %v", path, got, want)
		}
	default:
		if got == nil {
			t.Errorf("lookupPath(%q) = nil, want non-nil map", path)
		}
	}
}
