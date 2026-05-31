//ff:func feature=migration type=test control=sequence
//ff:what TestSortedMapKeys — 맵 키를 정렬된 슬라이스로 반환
package migration

import (
	"testing"
)

func TestSortedMapKeys(t *testing.T) {
	m := map[string]int{"c": 3, "a": 1, "b": 2}
	got := sortedMapKeys(m)
	if !stringSliceEqual(got, []string{"a", "b", "c"}) {
		t.Errorf("sortedMapKeys = %v, want [a b c]", got)
	}
	if len(sortedMapKeys(map[string]int{})) != 0 {
		t.Error("empty map should yield empty slice")
	}
}
