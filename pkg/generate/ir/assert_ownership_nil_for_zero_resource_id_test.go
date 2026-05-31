//ff:func feature=gen-ir type=test-helper control=sequence
//ff:what assertOwnershipNilForZeroResourceID — zero ResourceID 시 AuthOp.Ownership 이 nil 인지 검증 헬퍼
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// assertOwnershipNilForZeroResourceID builds a service plan whose @auth has a
// zero-valued ResourceID and asserts that Ownership is not populated.
func assertOwnershipNilForZeroResourceID(t *testing.T, fs *yongol.Fullstack, zero string) {
	t.Helper()
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
}
