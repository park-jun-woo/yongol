//ff:func feature=gen-ir type=test control=iteration dimension=1
//ff:what TestBuildServicePlanComplex -- 복합 워크플로우 IR 변환 (@get->@empty->@auth->@state->@put->@get->@response)

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanComplex(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Feature:  "workflow",
		Sequences: []ssac.Sequence{
			{Type: ssac.SeqGet, Model: "Workflow.FindByID", Inputs: map[string]string{"ID": "request.id"}, Result: &ssac.Result{Var: "wf", Type: "Workflow"}},
			{Type: ssac.SeqEmpty, Target: "wf", Message: "Workflow not found", ErrStatus: 404},
			{Type: ssac.SeqAuth, Action: "ArchiveWorkflow", Resource: "workflow", Inputs: map[string]string{"ResourceID": "wf.ID"}, Message: "Forbidden", ErrStatus: 403},
			{Type: ssac.SeqState, DiagramID: "Workflow", Inputs: map[string]string{"Status": "wf.Status"}, Transition: "ArchiveWorkflow", Message: "Cannot archive", ErrStatus: 409},
			{Type: ssac.SeqPut, Model: "Workflow.UpdateStatus", Inputs: map[string]string{"ID": "wf.ID", "Status": `"archived"`}},
			{Type: ssac.SeqGet, Model: "Workflow.FindByID", Inputs: map[string]string{"ID": "wf.ID"}, Result: &ssac.Result{Var: "updated", Type: "Workflow"}},
			{Type: ssac.SeqResponse, Fields: map[string]string{"workflow": "updated"}},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if plan.OperationID != "ArchiveWorkflow" {
		t.Errorf("OperationID = %q, want ArchiveWorkflow", plan.OperationID)
	}
	if !plan.UsesTransaction {
		t.Error("UsesTransaction = false, want true (has @put)")
	}
	if len(plan.Ops) != 7 {
		t.Fatalf("len(Ops) = %d, want 7", len(plan.Ops))
	}

	expected := []OpKind{OpGet, OpEmpty, OpAuth, OpState, OpPut, OpGet, OpResponse}
	for i, want := range expected {
		if plan.Ops[i].Kind != want {
			t.Errorf("Ops[%d].Kind = %d, want %d", i, plan.Ops[i].Kind, want)
		}
	}
	if plan.Ops[0].Get.FollowedByGuard != OpEmpty {
		t.Errorf("Ops[0].Get.FollowedByGuard = %d, want OpEmpty(%d)", plan.Ops[0].Get.FollowedByGuard, OpEmpty)
	}
	if plan.Ops[5].Get.FollowedByGuard != OpGet {
		t.Errorf("Ops[5].Get.FollowedByGuard = %d, want OpGet(0)", plan.Ops[5].Get.FollowedByGuard)
	}
	if len(plan.QueryMethods) < 2 {
		t.Errorf("len(QueryMethods) = %d, want >= 2", len(plan.QueryMethods))
	}
}
