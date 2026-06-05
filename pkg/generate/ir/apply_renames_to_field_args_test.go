//ff:func feature=gen-ir type=test control=sequence
//ff:what TestApplyRenamesToFieldArgs -- FieldArg.Source rename 맵 치환 검증
package ir

import "testing"

func TestApplyRenamesToFieldArgs(t *testing.T) {
	args := []FieldArg{
		{Source: "old"},
		{Source: "keep"},
		{Source: "old"},
	}
	renames := map[string]string{"old": "new"}
	applyRenamesToFieldArgs(args, renames)

	if args[0].Source != "new" || args[2].Source != "new" {
		t.Errorf("renamed sources = %q, %q, want new, new", args[0].Source, args[2].Source)
	}
	if args[1].Source != "keep" {
		t.Errorf("unmapped source = %q, want keep", args[1].Source)
	}
}
