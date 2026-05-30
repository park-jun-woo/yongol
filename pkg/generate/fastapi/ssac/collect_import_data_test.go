//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestCollectImportData — 여러 ServicePlan 스캔 후 import 메타데이터 집계

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectImportData(t *testing.T) {
	t.Run("EmptyPlans", func(t *testing.T) {
		d := collectImportData(nil, "feat")
		if d.Models == nil || d.ExtPkgs == nil {
			t.Fatal("expected initialized maps")
		}
		if len(d.Models) != 0 || d.HasAuth {
			t.Errorf("expected empty data, got %+v", d)
		}
	})

	t.Run("AggregatesAcrossPlans", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{Ops: []ir.Op{{
				Kind: ir.OpAuth,
				Auth: &ir.AuthOp{
					Action:   "UpdateWorkflow",
					Resource: "workflow",
					Ownership: &ir.OwnershipInfo{
						Table:       "workflows",
						OwnerColumn: "org_id",
						ResourcePK:  "id",
					},
				},
			}}},
			{Ops: []ir.Op{{Kind: ir.OpAuth, Auth: &ir.AuthOp{Action: "ListWorkflows", Resource: "workflow"}}}},
		}
		d := collectImportData(plans, "feat")
		if !d.HasAuth {
			t.Error("expected HasAuth true")
		}
		if !d.Models["Workflow"] {
			t.Errorf("expected Workflow model, got %v", d.Models)
		}
	})
}
