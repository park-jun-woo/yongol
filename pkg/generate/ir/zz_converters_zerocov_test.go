//ff:func feature=gen-ir type=test control=sequence
//ff:what TestConvertersZeroCov — 각 SSaC 시퀀스 종류를 BuildServicePlan 으로 변환해 convert* 디스패처 전체 커버

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// buildPlanOrFail runs BuildServicePlan with the given sequences and fails on error.
func buildPlanOrFail(t *testing.T, name string, seqs []ssac.Sequence) *ServicePlan {
	t.Helper()
	sf := &ssac.ServiceFunc{
		Name:      name,
		FileName:  name + ".ssac",
		Feature:   "feature",
		Sequences: seqs,
	}
	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("%s: BuildServicePlan error: %v", name, err)
	}
	if plan == nil {
		t.Fatalf("%s: nil plan", name)
	}
	return plan
}

func TestConvertGet_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "GetCourse", []ssac.Sequence{
		{Type: ssac.SeqGet, Model: "Course.FindByID",
			Inputs: map[string]string{"ID": "request.id"},
			Result: &ssac.Result{Var: "course", Type: "Course"}},
		{Type: ssac.SeqResponse, Target: "course"},
	})
	if plan.Ops[0].Kind != OpGet {
		t.Errorf("expected OpGet, got %d", plan.Ops[0].Kind)
	}
}

func TestConvertPost_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "CreateCourse", []ssac.Sequence{
		{Type: ssac.SeqPost, Model: "Course.Create",
			Inputs: map[string]string{"Title": "request.title", "InstructorID": "currentUser.ID"},
			Result: &ssac.Result{Var: "course", Type: "Course"}},
		{Type: ssac.SeqResponse, Fields: map[string]string{"course": "course"}},
	})
	if plan.Ops[0].Kind != OpPost {
		t.Errorf("expected OpPost, got %d", plan.Ops[0].Kind)
	}
}

func TestConvertPutDelete_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "ArchiveWorkflow", []ssac.Sequence{
		{Type: ssac.SeqPut, Model: "Workflow.UpdateStatus",
			Inputs: map[string]string{"ID": "request.id", "Status": `"archived"`}},
		{Type: ssac.SeqDelete, Model: "Workflow.Delete",
			Inputs: map[string]string{"ID": "request.id"}},
	})
	if plan.Ops[0].Kind != OpPut || plan.Ops[1].Kind != OpDelete {
		t.Errorf("expected OpPut/OpDelete, got %d/%d", plan.Ops[0].Kind, plan.Ops[1].Kind)
	}
}

func TestConvertAuth_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "ArchiveWorkflowAuth", []ssac.Sequence{
		{Type: ssac.SeqAuth, Action: "ArchiveWorkflow", Resource: "workflow",
			Inputs: map[string]string{"ResourceID": "wf.ID"}, Message: "Forbidden", ErrStatus: 403},
	})
	if plan.Ops[0].Kind != OpAuth {
		t.Errorf("expected OpAuth, got %d", plan.Ops[0].Kind)
	}
}

func TestConvertState_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "ArchiveWorkflowState", []ssac.Sequence{
		{Type: ssac.SeqState, DiagramID: "Workflow",
			Inputs: map[string]string{"Status": "wf.Status"},
			Transition: "ArchiveWorkflow", Message: "Cannot archive", ErrStatus: 409},
	})
	if plan.Ops[0].Kind != OpState {
		t.Errorf("expected OpState, got %d", plan.Ops[0].Kind)
	}
}

func TestConvertCall_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "RefreshToken", []ssac.Sequence{
		{Type: ssac.SeqCall, Model: "auth.RefreshRotate",
			Inputs: map[string]string{"RefreshToken": "request.refresh_token"},
			Result: &ssac.Result{Var: "rotated", Type: "auth.RefreshRotateResponse"}},
		{Type: ssac.SeqResponse, Fields: map[string]string{
			"access_token": "rotated.AccessToken", "refresh_token": "rotated.RefreshToken"}},
	})
	if plan.Ops[0].Kind != OpCall {
		t.Errorf("expected OpCall, got %d", plan.Ops[0].Kind)
	}
}

func TestConvertEval_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "CheckSchedule", []ssac.Sequence{
		{Type: ssac.SeqEval, Model: "schedule.IsExpired",
			Inputs: map[string]string{"StartsAt": "request.starts_at"},
			Message: "Schedule date is in the past", ErrStatus: 400},
	})
	if plan.Ops[0].Kind != OpEval {
		t.Errorf("expected OpEval, got %d", plan.Ops[0].Kind)
	}
}

func TestConvertPublish_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "CompleteOrder", []ssac.Sequence{
		{Type: ssac.SeqPost, Model: "Order.Complete",
			Inputs: map[string]string{"ID": "request.id"},
			Result: &ssac.Result{Var: "order", Type: "Order"}},
		{Type: ssac.SeqPublish, Topic: "order.completed",
			Inputs: map[string]string{"OrderID": "order.ID", "UserID": "currentUser.ID"}},
	})
	found := false
	for _, op := range plan.Ops {
		if op.Kind == OpPublish {
			found = true
		}
	}
	if !found {
		t.Errorf("expected an OpPublish in plan")
	}
}

func TestConvertVerifyPassword_ZeroCov(t *testing.T) {
	plan := buildPlanOrFail(t, "Login", []ssac.Sequence{
		{Type: ssac.SeqVerifyPassword, Model: "User",
			EmailCol: "email", EmailExpr: "request.body.email",
			HashCol: "password_hash", PasswordExpr: "request.body.password",
			Result: &ssac.Result{Var: "user", Type: "User"},
			ErrStatus: 401, Message: "Invalid credentials"},
		{Type: ssac.SeqResponse, Target: "user"},
	})
	if plan.Ops[0].Kind != OpVerifyPassword {
		t.Errorf("expected OpVerifyPassword, got %d", plan.Ops[0].Kind)
	}
}
