//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanPublish -- @publish 시퀀스 IR 변환 (topic/payload 검증)

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanPublish(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "CompleteOrder",
		FileName: "complete_order.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqPublish,
				Topic: "order.completed",
				Inputs: map[string]string{
					"OrderID": "order.ID",
					"UserID":  "currentUser.ID",
				},
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

	pubOp := plan.Ops[0]
	if pubOp.Kind != OpPublish {
		t.Fatalf("Ops[0].Kind = %d, want OpPublish", pubOp.Kind)
	}
	if pubOp.Publish.Topic != "order.completed" {
		t.Errorf("Publish.Topic = %q, want %q", pubOp.Publish.Topic, "order.completed")
	}
	if len(pubOp.Publish.Payload) != 2 {
		t.Fatalf("len(Publish.Payload) = %d, want 2", len(pubOp.Publish.Payload))
	}
}
