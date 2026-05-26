//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanSubscribe -- @subscribe 트리거 IR 변환 (TriggerSubscribe + Topic)

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanSubscribe(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "OnWorkflowExecuted",
		FileName: "on_workflow_executed.ssac",
		Subscribe: &ssac.SubscribeInfo{
			Topic: "workflow.executed",
		},
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqCall,
				Model: "webhook.Deliver",
				Inputs: map[string]string{
					"OrgID":      "message.OrgID",
					"WorkflowID": "message.WorkflowID",
					"Status":     "message.Status",
				},
				Result: &ssac.Result{Var: "_", Type: "webhook.DeliverResponse"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if plan.TriggerKind != TriggerSubscribe {
		t.Errorf("TriggerKind = %q, want %q", plan.TriggerKind, TriggerSubscribe)
	}
	if plan.Topic != "workflow.executed" {
		t.Errorf("Topic = %q, want %q", plan.Topic, "workflow.executed")
	}
	if len(plan.Ops) != 1 {
		t.Fatalf("len(Ops) = %d, want 1", len(plan.Ops))
	}
	if plan.Ops[0].Kind != OpCall {
		t.Errorf("Ops[0].Kind = %d, want OpCall", plan.Ops[0].Kind)
	}
}
