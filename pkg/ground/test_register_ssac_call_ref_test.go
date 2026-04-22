//ff:func feature=rule type=test control=sequence dimension=1
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

func TestRegisterSSaCCallRef_AlreadyCamel(t *testing.T) {
	set := map[string]bool{}
	registerSSaCCallRef("billing.charge", set)
	if !set["billing.charge"] {
		t.Errorf("expected billing.charge passthrough, got %v", set)
	}
}

func TestRegisterSSaCCallRef_NoDot(t *testing.T) {
	set := map[string]bool{}
	registerSSaCCallRef("NoPackage", set)
	if len(set) != 0 {
		t.Errorf("expected no insert when no dot, got %v", set)
	}
}

func TestRegisterSSaCCallRef_TrailingDot(t *testing.T) {
	set := map[string]bool{}
	registerSSaCCallRef("pkg.", set)
	if len(set) != 0 {
		t.Errorf("expected no insert when trailing dot (invalid), got %v", set)
	}
}
