//ff:func feature=gen-ir type=test control=sequence
//ff:what TestGetOpPaginationArgs -- GetOp.PaginationArgs 분리 검증 (cursor/per_page/limit/offset)
package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestGetOpAllPaginationKeys(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name:     "ListItems",
		FileName: "list_items.ssac",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqGet,
				Model: "Item.List",
				Inputs: map[string]string{
					"cursor":      "request.cursor",
					"per_page":    "request.per_page",
					"page_offset": "request.page_offset",
					"page":        "request.page",
					"limit":       "request.limit",
					"offset":      "request.offset",
					"Status":      "request.status",
				},
				Result: &ssac.Result{Var: "items", Type: "[]Item", Wrapper: "[]"},
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan: %v", err)
	}

	getOp := plan.Ops[0].Get
	if len(getOp.Args) != 1 {
		t.Errorf("len(Args) = %d, want 1 (Status only)", len(getOp.Args))
	}
	if len(getOp.PaginationArgs) != 6 {
		t.Errorf("len(PaginationArgs) = %d, want 6", len(getOp.PaginationArgs))
	}
}
