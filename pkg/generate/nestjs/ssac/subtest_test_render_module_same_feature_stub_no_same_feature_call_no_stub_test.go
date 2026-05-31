//ff:func feature=gen-nestjs type=test-helper control=sequence
//ff:what subtestTestRenderModuleSameFeatureStubNoSameFeatureCallNoStub — NoSameFeatureCallNoStub 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestRenderModuleSameFeatureStubNoSameFeatureCallNoStub(t *testing.T) {

	plans := []*ir.ServicePlan{
		{
			OperationID: "ListItems",
			HTTPMethod:  "GET",
			Feature:     "item",
			Ops: []ir.Op{
				{Kind: ir.OpGet, Get: &ir.GetOp{Model: "Item"}},
			},
		},
	}
	got, err := RenderModule("item", plans)
	if err != nil {
		t.Fatalf("RenderModule: %v", err)
	}

	// No stub service should be added.
	if strings.Contains(got, "ItemService,") {
		t.Errorf("should not add ItemService stub without same-feature @call, got:\n%s", got)
	}

}
