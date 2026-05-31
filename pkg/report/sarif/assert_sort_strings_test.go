//ff:func feature=report type=test-helper control=iteration dimension=1 topic=sarif
//ff:what assertSortStrings — sortStrings 결과가 want 와 일치하는지 검증하는 테스트 헬퍼
package sarif

import "testing"

// assertSortStrings copies in (preserving length) and asserts sortStrings
// produces want.
func assertSortStrings(t *testing.T, in, want []string) {
	t.Helper()
	// Copy in place without changing nil-ness vs empty-ness.
	s := make([]string, len(in))
	copy(s, in)
	sortStrings(s)
	if len(s) != len(want) {
		t.Fatalf("len: got %d, want %d", len(s), len(want))
	}
	for i := range want {
		if s[i] != want[i] {
			t.Errorf("index %d: got %q, want %q", i, s[i], want[i])
		}
	}
}
