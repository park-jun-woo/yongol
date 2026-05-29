//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanState -- @state 시퀀스 IR 변환 (diagram/transition/기본 상태코드 409)

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanState(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:       ssac.SeqState,
				DiagramID:  "Workflow",
				Inputs:     map[string]string{"Status": "wf.Status"},
				Transition: "ArchiveWorkflow",
				Message:    "Cannot archive",
				ErrStatus:  409,
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if len(plan.Ops) != 1 {
		t.Fatalf("len(Ops) = %d, want 1", len(plan.Ops))
	}

	stateOp := plan.Ops[0]
	if stateOp.Kind != OpState {
		t.Fatalf("Ops[0].Kind = %d, want OpState", stateOp.Kind)
	}
	if stateOp.State.Diagram != "Workflow" {
		t.Errorf("State.Diagram = %q, want %q", stateOp.State.Diagram, "Workflow")
	}
	if stateOp.State.Transition != "ArchiveWorkflow" {
		t.Errorf("State.Transition = %q, want %q", stateOp.State.Transition, "ArchiveWorkflow")
	}
	if stateOp.State.StatusCode != 409 {
		t.Errorf("State.StatusCode = %d, want 409", stateOp.State.StatusCode)
	}
	if len(stateOp.State.Inputs) != 1 {
		t.Fatalf("len(State.Inputs) = %d, want 1", len(stateOp.State.Inputs))
	}
	statusArg := stateOp.State.Inputs[0]
	if statusArg.Key != "Status" || statusArg.Source != "wf" || statusArg.Field != "Status" {
		t.Errorf("State.Inputs[0] = {Key:%q Source:%q Field:%q}, want {Status wf Status}", statusArg.Key, statusArg.Source, statusArg.Field)
	}
}
