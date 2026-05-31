//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestWriteHTTPHandlerEventBusEventBusAfterCurrentUser — EventBusAfterCurrentUser 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestWriteHTTPHandlerEventBusEventBusAfterCurrentUser(t *testing.T) {

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

}
