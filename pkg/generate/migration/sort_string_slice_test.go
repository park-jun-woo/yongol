//ff:func feature=migration type=test control=sequence
//ff:what TestSortStringSlice — 제자리 오름차순 정렬
package migration

import "testing"

func TestSortStringSlice(t *testing.T) {
	s := []string{"c", "a", "b"}
	sortStringSlice(s)
	if !stringSliceEqual(s, []string{"a", "b", "c"}) {
		t.Errorf("sortStringSlice produced %v, want [a b c]", s)
	}
}
