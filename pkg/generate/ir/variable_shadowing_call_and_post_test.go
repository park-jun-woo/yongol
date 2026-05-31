//ff:func feature=gen-ir type=test control=sequence
//ff:what TestVariableShadowing -- 동일 VarName 중복 선언 시 _result 접미사 자동 해소 + 후속 Op Source 갱신 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestVariableShadowingCallAndPost(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "MixedVars",
		FileName: "mixed.ssac",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqPost, Model: "A.Create",
				Result: &ssac.Result{Var: "result", Type: "A"}},
			{Type: ssac.SeqCall, Model: "pkg.Fn",
				Result: &ssac.Result{Var: "result", Type: "B"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	if plan.Ops[0].Post.VarName != "result" {
		t.Errorf("Post.VarName = %q, want result", plan.Ops[0].Post.VarName)
	}
	if plan.Ops[1].Call.ResultVar != "result_result" {
		t.Errorf("Call.ResultVar = %q, want result_result", plan.Ops[1].Call.ResultVar)
	}
}
