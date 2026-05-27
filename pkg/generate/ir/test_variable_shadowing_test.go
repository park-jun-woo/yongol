//ff:func feature=gen-ir type=test control=sequence
//ff:what TestVariableShadowing -- 동일 VarName 중복 선언 시 _result 접미사 자동 해소 검증

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
