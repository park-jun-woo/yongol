//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanGet -- @get 시퀀스 IR 변환 + 결과 바인딩 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanGet(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "GetCourse",
		FileName: "get_course.ssac",
		Feature:  "course",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Course.FindByID",
				Inputs: map[string]string{
					"ID": "request.id",
				},
				Result: &ssac.Result{
					Var:  "course",
					Type: "Course",
				},
			},
			{
				Type:   ssac.SeqResponse,
				Target: "course",
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if plan.OperationID != "GetCourse" {
		t.Errorf("OperationID = %q, want %q", plan.OperationID, "GetCourse")
	}
	if plan.TriggerKind != TriggerHTTP {
		t.Errorf("TriggerKind = %q, want %q", plan.TriggerKind, TriggerHTTP)
	}
	if plan.UsesTransaction {
		t.Error("UsesTransaction = true, want false for read-only plan")
	}
	if len(plan.Ops) != 2 {
		t.Fatalf("len(Ops) = %d, want 2", len(plan.Ops))
	}

	getOp := plan.Ops[0]
	if getOp.Kind != OpGet {
		t.Fatalf("Ops[0].Kind = %d, want OpGet", getOp.Kind)
	}
	if getOp.Get.VarName != "course" {
		t.Errorf("Get.VarName = %q, want %q", getOp.Get.VarName, "course")
	}
	if getOp.Get.VarType != "Course" {
		t.Errorf("Get.VarType = %q, want %q", getOp.Get.VarType, "Course")
	}
	if getOp.Get.Model != "Course" {
		t.Errorf("Get.Model = %q, want %q", getOp.Get.Model, "Course")
	}
	if getOp.Get.Method != "FindByID" {
		t.Errorf("Get.Method = %q, want %q", getOp.Get.Method, "FindByID")
	}
	if len(getOp.Get.Args) != 1 {
		t.Fatalf("len(Get.Args) = %d, want 1", len(getOp.Get.Args))
	}
	arg := getOp.Get.Args[0]
	if arg.Key != "ID" || arg.Source != "request" || arg.Field != "id" {
		t.Errorf("Get.Args[0] = {Key:%q Source:%q Field:%q}, want {ID request id}", arg.Key, arg.Source, arg.Field)
	}

	respOp := plan.Ops[1]
	if respOp.Kind != OpResponse {
		t.Fatalf("Ops[1].Kind = %d, want OpResponse", respOp.Kind)
	}
	if respOp.Response.SingleVar != "course" {
		t.Errorf("Response.SingleVar = %q, want %q", respOp.Response.SingleVar, "course")
	}
}
