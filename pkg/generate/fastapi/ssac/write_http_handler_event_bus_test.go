//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteHTTPHandlerEventBus — @publish 있는 핸들러에 event_bus DI 주입 + 전달 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteHTTPHandlerEventBus(t *testing.T) {
	t.Run("PublishOpInjectsEventBus", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "ExecuteWorkflow",
			HTTPMethod:  "POST",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/workflow/:id/execute",
			Feature:     "workflow",
			PathParams:  []string{"id"},
			Ops: []ir.Op{
				{Kind: ir.OpGet, Get: &ir.GetOp{Model: "Workflow"}},
				{Kind: ir.OpPublish, Publish: &ir.PublishOp{
					Topic: "workflow.executed",
				}},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()

		// Must inject EventBus dependency.
		if !strings.Contains(got, "event_bus: EventBus = Depends(get_event_bus)") {
			t.Errorf("expected EventBus dependency injection, got:\n%s", got)
		}
		// Must pass event_bus to service call.
		if !strings.Contains(got, "event_bus") {
			t.Errorf("expected event_bus in service call args, got:\n%s", got)
		}
		// Service call must include event_bus argument.
		svcCallIdx := strings.Index(got, "return await svc.execute_workflow(")
		if svcCallIdx < 0 {
			t.Fatalf("expected service call, got:\n%s", got)
		}
		svcCall := got[svcCallIdx:]
		if !strings.Contains(svcCall, "event_bus") {
			t.Errorf("service call must pass event_bus, got:\n%s", svcCall)
		}
	})

	t.Run("NoPublishOpSkipsEventBus", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "GetWorkflow",
			HTTPMethod:  "GET",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/workflow/:id",
			Feature:     "workflow",
			PathParams:  []string{"id"},
			Ops: []ir.Op{
				{Kind: ir.OpGet, Get: &ir.GetOp{Model: "Workflow"}},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()

		if strings.Contains(got, "EventBus") {
			t.Errorf("handler without @publish should not reference EventBus, got:\n%s", got)
		}
		if strings.Contains(got, "get_event_bus") {
			t.Errorf("handler without @publish should not reference get_event_bus, got:\n%s", got)
		}
	})

	t.Run("EventBusAfterCurrentUser", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "CompleteTask",
			HTTPMethod:  "PUT",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/task/:id/complete",
			Feature:     "task",
			PathParams:  []string{"id"},
			BodyFields: []ir.BodyFieldMeta{
				{Name: "note"},
			},
			Ops: []ir.Op{
				{Kind: ir.OpPut, Put: &ir.PutOp{Model: "Task"}},
				{Kind: ir.OpPublish, Publish: &ir.PublishOp{
					Topic: "task.completed",
				}},
			},
		}
		var b strings.Builder
		writeHTTPHandler(&b, plan)
		got := b.String()

		// event_bus should appear after session in params.
		sessionIdx := strings.Index(got, "Depends(get_session)")
		eventBusIdx := strings.Index(got, "Depends(get_event_bus)")
		if sessionIdx < 0 || eventBusIdx < 0 {
			t.Fatalf("expected both session and event_bus deps, got:\n%s", got)
		}
		if eventBusIdx < sessionIdx {
			t.Errorf("event_bus should come after session in params, got:\n%s", got)
		}

		// Service call should pass event_bus as last arg.
		svcCallIdx := strings.Index(got, "return await svc.complete_task(")
		if svcCallIdx < 0 {
			t.Fatalf("expected service call, got:\n%s", got)
		}
		svcCall := got[svcCallIdx:]
		if !strings.HasSuffix(strings.TrimSpace(svcCall), "event_bus)") {
			t.Errorf("event_bus should be last service call arg, got:\n%s", svcCall)
		}
	})
}
