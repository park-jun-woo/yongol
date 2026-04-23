//ff:func feature=rule type=test control=sequence
//ff:what registerSSaCCallRef — 후행 dot (잘못된 입력) 은 삽입되지 않음

package ground

import "testing"

func TestRegisterSSaCCallRef_TrailingDot(t *testing.T) {
	set := map[string]bool{}
	registerSSaCCallRef("pkg.", set)
	if len(set) != 0 {
		t.Errorf("expected no insert when trailing dot (invalid), got %v", set)
	}
}
