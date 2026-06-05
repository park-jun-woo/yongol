//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectModuleDeps — TestCollectModuleDeps — plan들에서 queue/authz/same-feature-stub + 정렬된 cross-feature @call 수집 검증

package ssac

import (
	"reflect"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectModuleDeps(t *testing.T) {
	plans := []*ir.ServicePlan{
		{Ops: []ir.Op{
			{Kind: ir.OpPublish},
			{Kind: ir.OpCall, Call: &ir.CallOp{TargetFeature: "shipping"}},
		}},
		{Ops: []ir.Op{
			{Kind: ir.OpAuth},
			{Kind: ir.OpCall, Call: &ir.CallOp{TargetFeature: "billing"}},
			{Kind: ir.OpCall, Call: &ir.CallOp{TargetFeature: "course"}}, // same feature -> stub
		}},
	}

	deps := collectModuleDeps("Course", plans)

	if !deps.NeedsQueue {
		t.Error("NeedsQueue should be true")
	}
	if !deps.NeedsAuthz {
		t.Error("NeedsAuthz should be true")
	}
	if !deps.NeedsSameFeatureStub {
		t.Error("NeedsSameFeatureStub should be true (call to same feature)")
	}
	// cross features sorted alphabetically
	if !reflect.DeepEqual(deps.CrossFeatures, []string{"billing", "shipping"}) {
		t.Errorf("CrossFeatures = %v, want [billing shipping]", deps.CrossFeatures)
	}
}
