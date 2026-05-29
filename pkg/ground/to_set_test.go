//ff:func feature=rule type=test control=sequence
//ff:what toSet — []string → rule.StringSet 변환 기본 케이스

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
