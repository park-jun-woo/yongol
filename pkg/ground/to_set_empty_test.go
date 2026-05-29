//ff:func feature=rule type=test control=sequence
//ff:what toSet — 빈/ nil 슬라이스 처리

package ground

import "testing"

func TestToSet_Empty(t *testing.T) {
	got := toSet(nil)
	if len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}
