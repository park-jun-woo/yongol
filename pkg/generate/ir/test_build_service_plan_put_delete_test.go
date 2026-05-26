//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanPutDelete -- @put/@delete 시퀀스 IR 변환 + 트랜잭션 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanPutDelete(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "DeleteWorkflow",
		FileName: "delete_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqPut,
				Model: "Workflow.UpdateStatus",
				Inputs: map[string]string{
					"ID":     "request.id",
					"Status": `"archived"`,
				},
			},
			{
				Type:  ssac.SeqDelete,
				Model: "Workflow.Delete",
				Inputs: map[string]string{
					"ID": "request.id",
				},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if !plan.UsesTransaction {
		t.Error("UsesTransaction = false, want true for plan with @put + @delete")
	}
	if len(plan.Ops) != 2 {
		t.Fatalf("len(Ops) = %d, want 2", len(plan.Ops))
	}

	putOp := plan.Ops[0]
	if putOp.Kind != OpPut {
		t.Fatalf("Ops[0].Kind = %d, want OpPut", putOp.Kind)
	}
	if putOp.Put.Model != "Workflow" || putOp.Put.Method != "UpdateStatus" {
		t.Errorf("Put = {%q %q}, want {Workflow UpdateStatus}", putOp.Put.Model, putOp.Put.Method)
	}
	// Check quoted literal arg.
	statusArg := findArgByKey(putOp.Put.Args, "Status")
	if statusArg == nil {
		t.Fatal("Put.Args missing Status key")
	}
	if statusArg.Literal != "archived" || !statusArg.IsQuoted {
		t.Errorf("Status arg = {Literal:%q IsQuoted:%v}, want {archived true}", statusArg.Literal, statusArg.IsQuoted)
	}

	delOp := plan.Ops[1]
	if delOp.Kind != OpDelete {
		t.Fatalf("Ops[1].Kind = %d, want OpDelete", delOp.Kind)
	}
	if delOp.Delete.Model != "Workflow" || delOp.Delete.Method != "Delete" {
		t.Errorf("Delete = {%q %q}, want {Workflow Delete}", delOp.Delete.Model, delOp.Delete.Method)
	}
}

