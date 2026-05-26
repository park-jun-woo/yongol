//ff:func feature=gen-ir type=test control=sequence
//ff:what TestAnnotateGetGuardsLastOp -- 마지막 Op이 @get → 가드 미주입 검증

package ir

import "testing"

func TestAnnotateGetGuardsLastOp(t *testing.T) {
	ops := []Op{
		{Kind: OpGet, Get: &GetOp{VarName: "wf"}},
	}
	annotateGetGuards(ops)
	if ops[0].Get.FollowedByGuard != OpGet {
		t.Errorf("FollowedByGuard = %d, want OpGet(0) for last-op-get", ops[0].Get.FollowedByGuard)
	}
}
