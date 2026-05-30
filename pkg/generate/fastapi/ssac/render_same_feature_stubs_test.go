//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderSameFeatureStubs — same-feature @call 대상 inline stub 블록 생성

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderSameFeatureStubs(t *testing.T) {
	t.Run("NoStubs", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{OperationID: "list_items", Ops: []ir.Op{{Kind: ir.OpGet, Get: &ir.GetOp{Model: "Item"}}}},
		}
		if got := RenderSameFeatureStubs(plans, "billing"); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
	t.Run("UndefinedCallTarget", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "checkout",
				Ops: []ir.Op{
					{Kind: ir.OpCall, Call: &ir.CallOp{Package: "billing", Function: "CalcTotal"}},
				},
			},
		}
		got := RenderSameFeatureStubs(plans, "billing")
		if !strings.Contains(got, "async def calc_total(*args, **kwargs):") {
			t.Errorf("expected calc_total stub, got %q", got)
		}
	})
}
