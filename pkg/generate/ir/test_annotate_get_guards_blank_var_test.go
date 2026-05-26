//ff:func feature=gen-ir type=test control=sequence
//ff:what TestAnnotateGetGuardsBlankVar -- 블랭크 변수("_") @get → 가드 미주입 검증

package ir

import "testing"

func TestAnnotateGetGuardsBlankVar(t *testing.T) {
	ops := []Op{
		{Kind: OpGet, Get: &GetOp{VarName: "_"}},
		{Kind: OpEmpty, Empty: &EmptyOp{VarName: "_"}},
	}
	annotateGetGuards(ops)
	if ops[0].Get.FollowedByGuard != OpGet {
		t.Errorf("FollowedByGuard = %d, want OpGet(0) for blank var", ops[0].Get.FollowedByGuard)
	}
}
