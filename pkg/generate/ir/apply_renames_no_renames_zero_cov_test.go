//ff:func feature=gen-ir type=test control=sequence
//ff:what TestVarShadowingZeroCov — resolveVariableShadowing / applyRenames / resolveVar 전 분기 직접 커버
package ir

import (
	"testing"
)

func TestApplyRenames_NoRenames_ZeroCov(t *testing.T) {
	op := &Op{Kind: OpGet, Get: &GetOp{Args: []FieldArg{{Source: "x"}}}}
	applyRenames(op, map[string]string{}) // empty → early return, no mutation
	if op.Get.Args[0].Source != "x" {
		t.Errorf("empty renames should not mutate")
	}
}
