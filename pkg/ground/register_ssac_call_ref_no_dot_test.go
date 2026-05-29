//ff:func feature=rule type=test control=sequence
//ff:what registerSSaCCallRef — dot 이 없는 입력은 삽입되지 않음

package ground

import "testing"

func TestRegisterSSaCCallRef_NoDot(t *testing.T) {
	set := map[string]bool{}
	registerSSaCCallRef("NoPackage", set)
	if len(set) != 0 {
		t.Errorf("expected no insert when no dot, got %v", set)
	}
}
