//ff:func feature=gen-ir type=test control=sequence
//ff:what TestVariableShadowing -- 동일 VarName 중복 선언 시 _result 접미사 자동 해소 + 후속 Op Source 갱신 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestVariableShadowingNoDuplicate(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "NoDup",
		FileName: "nodup.ssac",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqGet, Model: "A.X", Result: &ssac.Result{Var: "a", Type: "A"}},
			{Type: ssac.SeqGet, Model: "B.Y", Result: &ssac.Result{Var: "b", Type: "B"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	if plan.Ops[0].Get.VarName != "a" {
		t.Errorf("VarName[0] = %q, want a", plan.Ops[0].Get.VarName)
	}
	if plan.Ops[1].Get.VarName != "b" {
		t.Errorf("VarName[1] = %q, want b", plan.Ops[1].Get.VarName)
	}
}
