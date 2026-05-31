//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestWriteHTTPHandlerEventBusPublishOpInjectsEventBus — PublishOpInjectsEventBus 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestWriteHTTPHandlerEventBusPublishOpInjectsEventBus(t *testing.T) {

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

}
