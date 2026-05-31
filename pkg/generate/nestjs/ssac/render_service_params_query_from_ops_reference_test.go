//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_QueryFromOpsReference
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_QueryFromOpsReference(t *testing.T) {
	plan := &ir.ServicePlan{
		HTTPMethod:  "GET",
		QueryParams: nil,
		Ops: []ir.Op{
			{
				Kind: ir.OpGet,
				Get: &ir.GetOp{
					Args: []ir.FieldArg{
						{Key: "status", Location: ir.LocQuery},
					},
				},
			},
		},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "query: any") {
		t.Errorf("GET with LocQuery ops should include query, got %q", got)
	}
}
