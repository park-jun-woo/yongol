//ff:func feature=gen-nestjs type=test-helper control=sequence
//ff:what subtestTestRenderModuleSameFeatureStubNoCrossFeatureCallNoStub — NoCrossFeatureCallNoStub 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestRenderModuleSameFeatureStubNoCrossFeatureCallNoStub(t *testing.T) {

	plans := []*ir.ServicePlan{
		{
			OperationID: "CreateOrder",
			HTTPMethod:  "POST",
			Feature:     "order",
			Ops: []ir.Op{
				{Kind: ir.OpCall, Call: &ir.CallOp{
					Package:       "billing",
					TargetFeature: "billing",
					Function:      "HoldEscrow",
				}},
			},
		},
	}
	got, err := RenderModule("order", plans)
	if err != nil {
		t.Fatalf("RenderModule: %v", err)
	}

	// OrderService stub import should NOT be present (billing is cross-feature).
	if strings.Contains(got, "import { OrderService }") {
		t.Errorf("should not import OrderService for cross-feature call, got:\n%s", got)
	}

}
