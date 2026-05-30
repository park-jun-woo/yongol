//ff:func feature=gen-filefunc type=test control=sequence
//ff:what TestDefaultPatternEntries — 고정 pattern 키 맵 반환 검증

package filefunc

import "testing"

func TestDefaultPatternEntries(t *testing.T) {
	got := defaultPatternEntries()
	want := []string{"early-return", "error-collection"}
	if len(got) != len(want) {
		t.Errorf("expected %d patterns, got %d: %v", len(want), len(got), got)
	}
	for _, k := range want {
		if _, ok := got[k]; !ok {
			t.Errorf("missing pattern key %q: %v", k, got)
		}
	}
}
