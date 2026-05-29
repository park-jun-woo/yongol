//ff:func feature=rule type=test control=sequence
//ff:what registerSSaCModelRef — dot 이 없는 입력은 삽입되지 않음

package ground

import "testing"

// TestRegisterSSaCModelRef_NoDot — no dot → no insert.
func TestRegisterSSaCModelRef_NoDot(t *testing.T) {
	set := map[string]bool{}
	registerSSaCModelRef("NoDot", set)
	if len(set) != 0 {
		t.Errorf("expected empty set for no-dot input, got %v", set)
	}
}
