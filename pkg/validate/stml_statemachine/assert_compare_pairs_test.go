//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=stml-statemachine
//ff:what assertComparePairs — comparePair 슬라이스 두 개가 길이·순서·필드까지 동일한지 비교

package stml_statemachine

import "testing"

// assertComparePairs fails the test unless got equals want element-by-element,
// preserving order (DOM left-to-right is significant for collectComparePairs).
func assertComparePairs(t *testing.T, got, want []comparePair) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("pair count = %d, want %d (got %+v)", len(got), len(want), got)
	}
	for i := range want {
		if got[i] != want[i] {
			t.Errorf("pair[%d] = %+v, want %+v", i, got[i], want[i])
		}
	}
}
