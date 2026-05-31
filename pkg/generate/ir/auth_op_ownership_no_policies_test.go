//ff:func feature=gen-ir type=test control=sequence
//ff:what TestAuthOpOwnership -- AuthOp.Ownership Rego @ownership 이식 검증
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
