//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderSubscribeParams — renderSubscribeParams publish 유무에 따른 event_bus 추가 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderSubscribeParams(t *testing.T) {
	noPublish := &ir.ServicePlan{Ops: []ir.Op{{Kind: ir.OpCall}}}
	if got := renderSubscribeParams(noPublish); got != "session: AsyncSession, payload: dict" {
		t.Errorf("no publish: got %q", got)
	}

	withPublish := &ir.ServicePlan{Ops: []ir.Op{{Kind: ir.OpPublish}}}
	want := "session: AsyncSession, payload: dict, event_bus: EventBus | None = None"
	if got := renderSubscribeParams(withPublish); got != want {
		t.Errorf("with publish: got %q, want %q", got, want)
	}
}
