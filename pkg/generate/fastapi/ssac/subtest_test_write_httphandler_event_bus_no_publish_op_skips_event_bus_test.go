//ff:func feature=gen-fastapi type=test-helper control=sequence
//ff:what subtestTestWriteHTTPHandlerEventBusNoPublishOpSkipsEventBus — NoPublishOpSkipsEventBus 서브테스트
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func subtestTestWriteHTTPHandlerEventBusNoPublishOpSkipsEventBus(t *testing.T) {

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

}
