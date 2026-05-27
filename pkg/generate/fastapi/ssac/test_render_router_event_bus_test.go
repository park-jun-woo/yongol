//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderRouterEventBus — @publish plan 포함 시 event_bus import 주입 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderRouterEventBus(t *testing.T) {
	t.Run("PublishAddsEventBusImport", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "ExecuteWorkflow",
				HTTPMethod:  "POST",
				TriggerKind: ir.TriggerHTTP,
				URLPath:     "/workflow/:id/execute",
				Feature:     "workflow",
				PathParams:  []string{"id"},
				Ops: []ir.Op{
					{Kind: ir.OpPublish, Publish: &ir.PublishOp{
						Topic: "workflow.executed",
					}},
				},
			},
		}
		got, err := RenderRouter("workflow", plans)
		if err != nil {
			t.Fatalf("RenderRouter: %v", err)
		}

		if !strings.Contains(got, "from app.dependencies.event_bus import EventBus, get_event_bus") {
			t.Errorf("expected event_bus import, got:\n%s", got)
		}
	})

	t.Run("NoPublishSkipsEventBusImport", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "GetWorkflow",
				HTTPMethod:  "GET",
				TriggerKind: ir.TriggerHTTP,
				URLPath:     "/workflow/:id",
				Feature:     "workflow",
				PathParams:  []string{"id"},
				Ops: []ir.Op{
					{Kind: ir.OpGet, Get: &ir.GetOp{Model: "Workflow"}},
				},
			},
		}
		got, err := RenderRouter("workflow", plans)
		if err != nil {
			t.Fatalf("RenderRouter: %v", err)
		}

		if strings.Contains(got, "event_bus") {
			t.Errorf("should not contain event_bus import without @publish, got:\n%s", got)
		}
	})

	t.Run("MixedPlansAddImport", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "GetWorkflow",
				HTTPMethod:  "GET",
				TriggerKind: ir.TriggerHTTP,
				URLPath:     "/workflow/:id",
				Feature:     "workflow",
				PathParams:  []string{"id"},
				Ops: []ir.Op{
					{Kind: ir.OpGet, Get: &ir.GetOp{Model: "Workflow"}},
				},
			},
			{
				OperationID: "ExecuteWorkflow",
				HTTPMethod:  "POST",
				TriggerKind: ir.TriggerHTTP,
				URLPath:     "/workflow/:id/execute",
				Feature:     "workflow",
				PathParams:  []string{"id"},
				Ops: []ir.Op{
					{Kind: ir.OpPublish, Publish: &ir.PublishOp{
						Topic: "workflow.executed",
					}},
				},
			},
		}
		got, err := RenderRouter("workflow", plans)
		if err != nil {
			t.Fatalf("RenderRouter: %v", err)
		}

		if !strings.Contains(got, "from app.dependencies.event_bus import EventBus, get_event_bus") {
			t.Errorf("mixed plans with at least one @publish should add event_bus import, got:\n%s", got)
		}
	})
}
