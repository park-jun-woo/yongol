//ff:func feature=rule type=test control=sequence
//ff:what registerSSaCCallRef — 이미 camel 인 경우 그대로 passthrough

package ground

import "testing"

func TestRegisterSSaCCallRef_AlreadyCamel(t *testing.T) {
	set := map[string]bool{}
	registerSSaCCallRef("billing.charge", set)
	if !set["billing.charge"] {
		t.Errorf("expected billing.charge passthrough, got %v", set)
	}
}
