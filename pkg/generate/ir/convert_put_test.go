//ff:func feature=gen-ir type=test control=sequence
//ff:what convertDelete/convertPut/convertEmpty/convertExists/matchFollowingGuard/resolveVar/convertInputsToFieldArgs
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestConvertPut(t *testing.T) {
	op := convertPut(ssac.Sequence{Model: "Course.Update"})
	if op.Kind != OpPut || op.Put == nil || op.Put.Model != "Course" || op.Put.Method != "Update" {
		t.Errorf("op = %+v", op)
	}
}
