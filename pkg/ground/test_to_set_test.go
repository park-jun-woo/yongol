//ff:func feature=rule type=test control=sequence dimension=1
//ff:what toSet — []string → rule.StringSet 변환 + 빈 슬라이스/중복 처리

package ground

import "testing"

func TestToSet_Basic(t *testing.T) {
	got := toSet([]string{"a", "b", "c"})
	if !got["a"] || !got["b"] || !got["c"] {
		t.Fatalf("toSet missing entries: %v", got)
	}
	if len(got) != 3 {
		t.Errorf("len = %d, want 3", len(got))
	}
}

func TestToSet_Duplicates(t *testing.T) {
	got := toSet([]string{"a", "a", "b"})
	if len(got) != 2 {
		t.Errorf("len = %d, want 2 (dedup)", len(got))
	}
}

func TestToSet_Empty(t *testing.T) {
	got := toSet(nil)
	if len(got) != 0 {
		t.Errorf("len = %d, want 0", len(got))
	}
}
