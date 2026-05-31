//ff:func feature=gen-gogin type=test control=sequence
//ff:what appendQueueSubscribeLines — ServiceFunc @subscribe에서 queue.Subscribe 라인 추가
package boot

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestAppendQueueSubscribeLines(t *testing.T) {
	funcs := []ssac.ServiceFunc{
		{Name: "OnOrderCompleted", Subscribe: &ssac.SubscribeInfo{Topic: "order.completed"}},
		{Name: "HttpHandler"}, // no subscribe → skipped
		{Name: "OnCartAbandoned", Subscribe: &ssac.SubscribeInfo{Topic: "cart.abandoned"}},
	}
	out := appendQueueSubscribeLines([]string{"prefix"}, funcs)
	body := strings.Join(out, "\n")

	if out[0] != "prefix" {
		t.Errorf("existing lines must be preserved, got %q", out[0])
	}
	if !strings.Contains(body, `queue.Subscribe("order.completed", srv.OnOrderCompleted)`) {
		t.Errorf("missing order subscribe line, got:\n%s", body)
	}
	if !strings.Contains(body, `queue.Subscribe("cart.abandoned", srv.OnCartAbandoned)`) {
		t.Errorf("missing cart subscribe line, got:\n%s", body)
	}
	if strings.Contains(body, "HttpHandler") {
		t.Errorf("non-subscribe func must be skipped, got:\n%s", body)
	}
	// 1 prefix + 2 subscribe lines.
	if len(out) != 3 {
		t.Errorf("expected 3 lines, got %d: %v", len(out), out)
	}
}
