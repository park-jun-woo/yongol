//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteRouterHandlers — writeRouterHandlers HTTP/subscribe 트리거별 핸들러 분기 출력 검증
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteRouterHandlers(t *testing.T) {
	plans := []*ir.ServicePlan{
		{
			OperationID: "GetWorkflow",
			HTTPMethod:  "GET",
			TriggerKind: ir.TriggerHTTP,
			URLPath:     "/workflow/:id",
			Feature:     "workflow",
			PathParams:  []string{"id"},
		},
		{
			OperationID: "OnOrderPaid",
			TriggerKind: ir.TriggerSubscribe,
			Topic:       "order.paid",
		},
	}
	var b strings.Builder
	writeRouterHandlers(&b, plans)
	got := b.String()

	if !strings.Contains(got, "Depends(get_session)") {
		t.Errorf("expected HTTP handler output, got:\n%s", got)
	}
	if !strings.Contains(got, "Subscribe handler for topic: order.paid") {
		t.Errorf("expected subscribe handler output, got:\n%s", got)
	}
}
