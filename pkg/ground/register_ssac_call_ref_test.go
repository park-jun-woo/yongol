//ff:func feature=rule type=test control=sequence
//ff:what registerSSaCCallRef — "pkg.Func" → "pkg.func" 정규화

package ground

import "testing"

func TestRegisterSSaCCallRef_PascalToCamel(t *testing.T) {
	set := map[string]bool{}
	registerSSaCCallRef("auth.HashPassword", set)
	if !set["auth.hashPassword"] {
		t.Errorf("expected auth.hashPassword in set, got %v", set)
	}
}
