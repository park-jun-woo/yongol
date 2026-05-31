//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_PathFromOpsReference
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_PathFromOpsReference(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod: "GET",
		PathParams: nil,
		Ops: []ir.Op{
			{
				Kind: ir.OpGet,
				Get: &ir.GetOp{
					Args: []ir.FieldArg{
						{Key: "id", Location: ir.LocPath},
					},
				},
			},
		},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "params: any") {
		t.Errorf("GET with LocPath ops should include params, got %q", got)
	}
}
