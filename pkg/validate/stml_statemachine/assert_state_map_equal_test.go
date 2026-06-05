//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-statemachine
//ff:what assertStateMapEqual — Symbol→상태집합 맵 두 개가 키·집합까지 동일한지 비교

package stml_statemachine

import "testing"

// assertStateMapEqual fails the test unless got equals want, comparing both the
// set of Symbol keys and the inner state-name sets exactly.
func assertStateMapEqual(t *testing.T, got, want map[string]map[string]bool) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("key count = %d, want %d (got %+v)", len(got), len(want), got)
	}
	for sym, wantSet := range want {
		gotSet, ok := got[sym]
		if !ok || len(gotSet) != len(wantSet) {
			t.Fatalf("symbol %q set = %+v, want %+v", sym, gotSet, wantSet)
		}
		assertStateSetEqual(t, sym, gotSet, wantSet)
	}
}
