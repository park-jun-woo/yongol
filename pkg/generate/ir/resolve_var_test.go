//ff:func feature=gen-ir type=test control=sequence
//ff:what convertDelete/convertPut/convertEmpty/convertExists/matchFollowingGuard/resolveVar/convertInputsToFieldArgs
package ir

import (
	"testing"
)

func TestResolveVar(t *testing.T) {
	declared := map[string]bool{}
	if got := resolveVar("user", declared); got != "user" {
		t.Errorf("first use = %q, want user", got)
	}
	// second use collides -> _result suffix
	if got := resolveVar("user", declared); got != "user_result" {
		t.Errorf("second use = %q, want user_result", got)
	}
	// third use collides again
	if got := resolveVar("user", declared); got != "user_result_result" {
		t.Errorf("third use = %q, want user_result_result", got)
	}
}
