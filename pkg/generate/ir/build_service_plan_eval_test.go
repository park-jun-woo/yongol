//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanEval -- @eval 시퀀스 IR 변환 (bool 가드 + 기본 상태코드 400)

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanEval(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "SetSchedule",
		FileName: "set_schedule.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqEval,
				Model: "schedule.IsExpired",
				Inputs: map[string]string{
					"StartsAt": "request.starts_at",
				},
				Message:   "Schedule date is in the past",
				ErrStatus: 400,
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

	evalOp := plan.Ops[0]
	if evalOp.Kind != OpEval {
		t.Fatalf("Ops[0].Kind = %d, want OpEval", evalOp.Kind)
	}
	if evalOp.Eval.Package != "schedule" {
		t.Errorf("Eval.Package = %q, want %q", evalOp.Eval.Package, "schedule")
	}
	if evalOp.Eval.Function != "IsExpired" {
		t.Errorf("Eval.Function = %q, want %q", evalOp.Eval.Function, "IsExpired")
	}
	if evalOp.Eval.StatusCode != 400 {
		t.Errorf("Eval.StatusCode = %d, want 400", evalOp.Eval.StatusCode)
	}
	if evalOp.Eval.Message != "Schedule date is in the past" {
		t.Errorf("Eval.Message = %q, want expected", evalOp.Eval.Message)
	}
}
