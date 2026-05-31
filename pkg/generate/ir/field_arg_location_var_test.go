//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgLocation -- FieldArg.Location 분류 검증 (path/query/body/var/literal/user)
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFieldArgLocationVar(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Workflow.FindByID",
				Inputs: map[string]string{
					"ID": "request.id",
				},
				Result: &ssac.Result{Var: "wf", Type: "Workflow"},
			},
			{
				Type:  ssac.SeqPut,
				Model: "Workflow.UpdateStatus",
				Inputs: map[string]string{
					"ID": "wf.ID",
				},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	putOp := plan.Ops[1]
	idArg := findArgByKey(putOp.Put.Args, "ID")
	if idArg == nil {
		t.Fatal("missing ID arg")
	}
	if idArg.Location != LocVar {
		t.Errorf("ID.Location = %q, want %q", idArg.Location, LocVar)
	}
}
