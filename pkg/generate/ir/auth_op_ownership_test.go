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

func TestAuthOpOwnership(t *testing.T) {
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
		Name:     "ArchiveWorkflow",
		FileName: "archive_workflow.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:     ssac.SeqAuth,
				Action:   "ArchiveWorkflow",
				Resource: "workflow",
				Inputs: map[string]string{
					"ResourceID": "wf.ID",
				},
				Message:   "Forbidden",
				ErrStatus: 403,
			},
		},
	}

	plan, err := BuildServicePlan(sf, fs)
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	authOp := plan.Ops[0]
	if authOp.Kind != OpAuth {
		t.Fatalf("Ops[0].Kind = %d, want OpAuth", authOp.Kind)
	}
	if authOp.Auth.Ownership == nil {
		t.Fatal("Auth.Ownership = nil, want non-nil")
	}
	ow := authOp.Auth.Ownership
	if ow.Table != "workflows" {
		t.Errorf("Ownership.Table = %q, want %q", ow.Table, "workflows")
	}
	if ow.OwnerColumn != "owner_id" {
		t.Errorf("Ownership.OwnerColumn = %q, want %q", ow.OwnerColumn, "owner_id")
	}
	if ow.ResourcePK != "id" {
		t.Errorf("Ownership.ResourcePK = %q, want %q", ow.ResourcePK, "id")
	}
}

func TestAuthOpOwnershipNoMapping(t *testing.T) {
	fs := &yongol.Fullstack{
		ParsedPolicies: []rego.Policy{
			{
				Ownerships: []rego.OwnershipMapping{
					{Resource: "project", Table: "projects", Column: "owner_id"},
				},
			},
		},
	}

	sf := &ssac.ServiceFunc{
		Name:     "DeleteCourse",
		FileName: "delete_course.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:     ssac.SeqAuth,
				Action:   "DeleteCourse",
				Resource: "course",
				Inputs:   map[string]string{"ResourceID": "c.ID"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, fs)
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}
	if plan.Ops[0].Auth.Ownership != nil {
		t.Errorf("Ownership = %+v, want nil for non-matching resource", plan.Ops[0].Auth.Ownership)
	}
}

// TestAuthOpOwnershipNoResourceID verifies that when Inputs has no ResourceID
// key (e.g. @auth {} "Forbidden"), Ownership is not populated even if a
// matching @ownership Rego annotation exists. This is the Group B fix.
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

// TestAuthOpOwnershipZeroResourceID verifies that a zero-value ResourceID
// ("0", "", "nil", "null") also skips ownership enrichment.
func TestAuthOpOwnershipZeroResourceID(t *testing.T) {
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

	for _, zero := range []string{"0", "", "nil", "null"} {
		t.Run(zero, func(t *testing.T) {
			sf := &ssac.ServiceFunc{
				Name:     "ListWorkflows",
				FileName: "list_workflows.ssac",
				Sequences: []ssac.Sequence{
					{
						Type:     ssac.SeqAuth,
						Action:   "ListWorkflows",
						Resource: "workflow",
						Inputs:   map[string]string{"ResourceID": zero},
						Message:  "Forbidden",
					},
				},
			}

			plan, err := BuildServicePlan(sf, fs)
			if err != nil {
				t.Fatalf("BuildServicePlan: %v", err)
			}
			if plan.Ops[0].Auth.Ownership != nil {
				t.Errorf("Ownership = %+v, want nil for zero ResourceID %q",
					plan.Ops[0].Auth.Ownership, zero)
			}
		})
	}
}

func TestAuthOpOwnershipNoPolicies(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "DeleteCourse",
		FileName: "delete_course.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:     ssac.SeqAuth,
				Action:   "DeleteCourse",
				Resource: "course",
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}
	if plan.Ops[0].Auth.Ownership != nil {
		t.Error("Ownership should be nil when no policies")
	}
}
