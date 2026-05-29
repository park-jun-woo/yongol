//ff:func feature=rule type=test control=sequence
//ff:what toSet — 중복 입력은 dedup

package ground

import "testing"

func TestToSet_Duplicates(t *testing.T) {
	got := toSet([]string{"a", "a", "b"})
	if len(got) != 2 {
		t.Errorf("len = %d, want 2 (dedup)", len(got))
	}
}
