//ff:func feature=gen-ir type=test control=sequence
//ff:what TestAuthOpOwnership -- AuthOp.Ownership Rego @ownership 이식 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ddl"
	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestAuthOpOwnershipNoResourceID(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{
				Ownerships: []rego.OwnershipMapping{
					{Resource: "workflow", Table: "workflows", Column: "owner_id"},
				},
			},
		},
		DDLTables: []ddl.Table{
			{
				Name: "workflows",
				Columns: map[string]ddl.Column{
					"id":       {Name: "id"},
					"owner_id": {Name: "owner_id"},
				},
				PrimaryKey: []string{"id"},
			},
		},
	}

	sf := &ssac.ServiceFunc{
		Name:     "CreateWorkflow",
		FileName: "create_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:     ssac.SeqAuth,
				Action:   "CreateWorkflow",
				Resource: "workflow",
				Inputs:   map[string]string{},
				Message:  "Forbidden",
			},
		},
	}

	plan, err := BuildServicePlan(sf, fs)
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}
	if plan.Ops[0].Auth.Ownership != nil {
		t.Errorf("Ownership = %+v, want nil when no ResourceID in Inputs",
			plan.Ops[0].Auth.Ownership)
	}
}
