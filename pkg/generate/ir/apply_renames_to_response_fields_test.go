//ff:func feature=gen-ir type=test control=sequence
//ff:what TestApplyRenamesToResponseFields -- ResponseField.Source dot-notation rename 치환 검증
package ir

import "testing"

func TestApplyRenamesToResponseFields(t *testing.T) {
	fields := []ResponseField{
		{Source: "user"},
		{Source: "user.email"},
		{Source: "keep.id"},
	}
	renames := map[string]string{"user": "u"}
	applyRenamesToResponseFields(fields, renames)

	if fields[0].Source != "u" {
		t.Errorf("plain source = %q, want u", fields[0].Source)
	}
	if fields[1].Source != "u.email" {
		t.Errorf("dotted source = %q, want u.email", fields[1].Source)
	}
	if fields[2].Source != "keep.id" {
		t.Errorf("unmapped dotted source = %q, want keep.id", fields[2].Source)
	}
}
