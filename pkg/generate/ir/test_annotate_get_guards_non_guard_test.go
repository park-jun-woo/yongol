//ff:func feature=gen-ir type=test control=sequence
//ff:what TestAnnotateGetGuardsNonGuardFollower -- 비가드 후속(@auth) → 가드 미주입 검증

package ir

import "testing"

func TestAnnotateGetGuardsNonGuardFollower(t *testing.T) {
	ops := []Op{
		{Kind: OpGet, Get: &GetOp{VarName: "wf"}},
		{Kind: OpAuth, Auth: &AuthOp{Action: "Delete"}},
	}
	annotateGetGuards(ops)
	if ops[0].Get.FollowedByGuard != OpGet {
		t.Errorf("FollowedByGuard = %d, want OpGet(0) for non-guard follower", ops[0].Get.FollowedByGuard)
	}
}
