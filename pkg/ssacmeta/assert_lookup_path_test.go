//ff:func feature=ssacmeta type=test-helper control=sequence
//ff:what assertLookupPath — lookupPath 결과(ok + scalar/map 값) 검증 헬퍼
package ssacmeta

import "testing"

// assertLookupPath asserts lookupPath(m, path) ok-ness and value. Scalars are
// compared directly; map values are only checked for non-nil.
func assertLookupPath(t *testing.T, m map[string]any, path string, want any, wantOk bool) {
	t.Helper()
	got, ok := lookupPath(m, path)
	if ok != wantOk {
		t.Fatalf("lookupPath(%q) ok = %v, want %v", path, ok, wantOk)
	}
	if !ok {
		return
	}
	assertLookupValue(t, path, got, want)
}
