//ff:func feature=gen-ir type=test control=sequence
//ff:what TestAnnotateGetGuardsNoMatch -- 다른 변수명 @empty → 가드 미주입 검증

package ir

import "testing"

func TestAnnotateGetGuardsNoMatch(t *testing.T) {
	ops := []Op{
		{Kind: OpGet, Get: &GetOp{VarName: "wf"}},
		{Kind: OpEmpty, Empty: &EmptyOp{VarName: "other"}},
	}
	annotateGetGuards(ops)
	if ops[0].Get.FollowedByGuard != OpGet {
		t.Errorf("FollowedByGuard = %d, want OpGet(0) for mismatched var", ops[0].Get.FollowedByGuard)
	}
}
