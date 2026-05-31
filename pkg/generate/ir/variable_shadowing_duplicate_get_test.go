//ff:func feature=gen-ir type=test control=sequence
//ff:what TestVariableShadowing -- 동일 VarName 중복 선언 시 _result 접미사 자동 해소 + 후속 Op Source 갱신 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestVariableShadowingDuplicateGet(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "DuplicateVar",
		FileName: "dup.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:   ssac.SeqGet,
				Model:  "Course.FindByID",
				Inputs: map[string]string{"ID": "request.id"},
				Result: &ssac.Result{Var: "item", Type: "Course"},
			},
			{
				Type:   ssac.SeqGet,
				Model:  "Course.FindBySlug",
				Inputs: map[string]string{"Slug": "request.slug"},
				Result: &ssac.Result{Var: "item", Type: "Course"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	if plan.Ops[0].Get.VarName != "item" {
		t.Errorf("Ops[0].Get.VarName = %q, want %q", plan.Ops[0].Get.VarName, "item")
	}
	if plan.Ops[1].Get.VarName != "item_result" {
		t.Errorf("Ops[1].Get.VarName = %q, want %q", plan.Ops[1].Get.VarName, "item_result")
	}
}
