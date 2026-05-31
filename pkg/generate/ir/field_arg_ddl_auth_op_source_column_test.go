//ff:func feature=gen-ir type=test control=sequence
//ff:what TestFieldArgDDL -- FieldArg.ColumnName/IsPK DDL 매핑 + @call/@auth SourceColumn 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestFieldArgDDLAuthOpSourceColumn(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "DeleteWorkflow",
		FileName: "delete_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:   ssac.SeqGet,
				Model:  "Workflow.FindByID",
				Inputs: map[string]string{"ID": "request.ID"},
				Result: &ssac.Result{Var: "wf", Type: "Workflow"},
			},
			{
				Type:   ssac.SeqAuth,
				Model:  "workflow.delete",
				Inputs: map[string]string{"ResourceID": "wf.OrgID"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	authOp := plan.Ops[1]
	if authOp.Kind != OpAuth {
		t.Fatalf("Ops[1].Kind = %d, want OpAuth", authOp.Kind)
	}

	resArg := findArgByKey(authOp.Auth.Inputs, "ResourceID")
	if resArg == nil {
		t.Fatal("missing ResourceID arg")
	}
	if resArg.SourceColumn != "org_id" {
		t.Errorf("ResourceID.SourceColumn = %q, want %q", resArg.SourceColumn, "org_id")
	}
	if resArg.Source != "wf" {
		t.Errorf("ResourceID.Source = %q, want %q", resArg.Source, "wf")
	}
}
