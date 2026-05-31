//ff:func feature=gen-ir type=test control=sequence
//ff:what TestVarShadowingZeroCov — resolveVariableShadowing / applyRenames / resolveVar 전 분기 직접 커버
package ir

import (
	"testing"
)

func TestResolveVar_ZeroCov(t *testing.T) {
	declared := map[string]bool{}
	if got := resolveVar("user", declared); got != "user" {
		t.Errorf("first use = %q, want user", got)
	}
	if got := resolveVar("user", declared); got != "user_result" {
		t.Errorf("collision = %q, want user_result", got)
	}
	if got := resolveVar("user", declared); got != "user_result_result" {
		t.Errorf("double collision = %q, want user_result_result", got)
	}
}
