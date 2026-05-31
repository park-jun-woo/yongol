//ff:func feature=gen-ir type=test control=sequence
//ff:what convertDelete/convertPut/convertEmpty/convertExists/matchFollowingGuard/resolveVar/convertInputsToFieldArgs
package ir

import (
	"testing"
)

func TestMatchFollowingGuard(t *testing.T) {
	empty := Op{Kind: OpEmpty, Empty: &EmptyOp{VarName: "course"}}
	if got := matchFollowingGuard(empty, "course"); got != OpEmpty {
		t.Errorf("empty match = %v, want OpEmpty", got)
	}
	exists := Op{Kind: OpExists, Exists: &ExistsOp{VarName: "dup"}}
	if got := matchFollowingGuard(exists, "dup"); got != OpExists {
		t.Errorf("exists match = %v, want OpExists", got)
	}
	// non-matching var name -> OpGet (zero)
	if got := matchFollowingGuard(empty, "other"); got != OpGet {
		t.Errorf("non-match = %v, want OpGet", got)
	}
	// non-guard op -> OpGet
	if got := matchFollowingGuard(Op{Kind: OpPost}, "x"); got != OpGet {
		t.Errorf("non-guard = %v, want OpGet", got)
	}
}
