//ff:func feature=gen-ir type=test control=sequence
//ff:what TestGetOpPaginationArgs -- GetOp.PaginationArgs 분리 검증 (cursor/per_page/limit/offset)
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGetOpNoPaginationArgs(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "GetCourse",
		FileName: "get_course.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Course.FindByID",
				Inputs: map[string]string{
					"ID": "request.id",
				},
				Result: &ssac.Result{Var: "course", Type: "Course"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	getOp := plan.Ops[0].Get
	if len(getOp.Args) != 1 {
		t.Errorf("len(Args) = %d, want 1", len(getOp.Args))
	}
	if len(getOp.PaginationArgs) != 0 {
		t.Errorf("len(PaginationArgs) = %d, want 0", len(getOp.PaginationArgs))
	}
}
