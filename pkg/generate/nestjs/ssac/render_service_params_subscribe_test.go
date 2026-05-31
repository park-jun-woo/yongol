//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderServiceParams_Subscribe
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParams_Subscribe(t *testing.T) {
	plan := &ir.ServicePlan{TriggerKind: ir.TriggerSubscribe}
	got := renderServiceParams(plan)
	if got != "payload: any" {
		t.Errorf("subscribe params = %q, want %q", got, "payload: any")
	}
}
