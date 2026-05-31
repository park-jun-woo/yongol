//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_BodyFromOpsReference
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_BodyFromOpsReference(t *testing.T) {
	// POST with empty BodyFields but Ops referencing LocBody
	plan := &ir.ServicePlan{
		HTTPMethod: "POST",
		BodyFields: nil,
		Ops: []ir.Op{
			{
				Kind: ir.OpPost,
				Post: &ir.PostOp{
					Args: []ir.FieldArg{
						{Key: "title", Location: ir.LocBody},
					},
				},
			},
		},
	}
	got := renderServiceParams(plan)
	if !strings.Contains(got, "body: any") {
		t.Errorf("POST with LocBody ops should include body param, got %q", got)
	}
}
