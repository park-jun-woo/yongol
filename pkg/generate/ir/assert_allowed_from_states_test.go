//ff:func feature=gen-ir type=test-helper control=iteration dimension=1
//ff:what assertAllowedFromStates — lookupAllowedFromStates 결과 슬라이스를 want와 길이·원소 비교하는 헬퍼

package ir

import "testing"

// assertAllowedFromStates asserts got equals want element-wise.
func assertAllowedFromStates(t *testing.T, got, want []string) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("len = %d, want %d (got %v, want %v)", len(got), len(want), got, want)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("[%d] = %q, want %q", i, got[i], want[i])
		}
	}
}
