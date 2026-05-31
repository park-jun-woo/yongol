//ff:func feature=gen-ir type=test control=sequence
//ff:what TestAuthOpOwnership -- AuthOp.Ownership Rego @ownership 이식 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/rego"
	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
