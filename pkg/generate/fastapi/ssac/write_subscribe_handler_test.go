//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteSubscribeHandler — 큐 구독 핸들러 함수 작성

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteSubscribeHandler(t *testing.T) {
	plan := &ir.ServicePlan{
		OperationID: "OnOrderPaid",
		Topic:       "order.paid",
	}
	var b strings.Builder
	writeSubscribeHandler(&b, plan)
	got := b.String()
	want := "# Subscribe handler for topic: order.paid\n" +
		"async def handle_on_order_paid(session: AsyncSession, payload: dict):\n" +
		"    return await svc.on_order_paid(session, payload)\n"
	if got != want {
		t.Errorf("got:\n%q\nwant:\n%q", got, want)
	}
}
