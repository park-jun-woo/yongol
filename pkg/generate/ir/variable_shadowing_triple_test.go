//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestVariableShadowing -- 동일 VarName 중복 선언 시 _result 접미사 자동 해소 + 후속 Op Source 갱신 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestVariableShadowingTriple(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "TripleVar",
		FileName: "triple.ssac",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqGet, Model: "A.X", Result: &ssac.Result{Var: "v", Type: "A"}},
			{Type: ssac.SeqGet, Model: "B.Y", Result: &ssac.Result{Var: "v", Type: "B"}},
			{Type: ssac.SeqGet, Model: "C.Z", Result: &ssac.Result{Var: "v", Type: "C"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	names := []string{
		plan.Ops[0].Get.VarName,
		plan.Ops[1].Get.VarName,
		plan.Ops[2].Get.VarName,
	}
	want := []string{"v", "v_result", "v_result_result"}
	for i, n := range names {
		if n != want[i] {
			t.Errorf("Ops[%d].Get.VarName = %q, want %q", i, n, want[i])
		}
	}
}
