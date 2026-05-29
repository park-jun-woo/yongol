//ff:func feature=gen-ir type=test control=sequence
//ff:what TestBuildServicePlanResponseEmpty -- @response 빈 응답 (204 No Content) IR 검증

package ir

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestBuildServicePlanResponseEmpty(t *testing.T) {
	sf := &ssac.ServiceFunc{
		Name: "DeleteItem",
		Sequences: []ssac.Sequence{
			{
				Type:  ssac.SeqDelete,
				Model: "Item.Delete",
				Inputs: map[string]string{
					"ID": "request.id",
				},
			},
			{
				Type: ssac.SeqResponse,
			},
		},
	}

	plan, err := BuildServicePlan(sf, &yongol.Fullstack{})
	if err != nil {
		t.Fatalf("BuildServicePlan returned error: %v", err)
	}

	if len(plan.Ops) != 2 {
		t.Fatalf("len(Ops) = %d, want 2", len(plan.Ops))
	}

	respOp := plan.Ops[1]
	if respOp.Kind != OpResponse {
		t.Fatalf("Ops[1].Kind = %d, want OpResponse", respOp.Kind)
	}
	if respOp.Response.SingleVar != "" {
		t.Errorf("Response.SingleVar = %q, want empty", respOp.Response.SingleVar)
	}
	if len(respOp.Response.Fields) != 0 {
		t.Errorf("len(Response.Fields) = %d, want 0", len(respOp.Response.Fields))
	}
}
