//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRouterDependencyFlags — routerDependencyFlags 인증/event_bus 의존성 판정 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRouterDependencyFlags(t *testing.T) {
	t.Run("HTTPNeedsAuth", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{TriggerKind: ir.TriggerHTTP, Ops: []ir.Op{{Kind: ir.OpCall}}},
		}
		auth, bus := routerDependencyFlags(plans)
		if !auth || bus {
			t.Errorf("got auth=%v bus=%v, want true,false", auth, bus)
		}
	})

	t.Run("VerifyPasswordPreAuth", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{TriggerKind: ir.TriggerHTTP, Ops: []ir.Op{{Kind: ir.OpVerifyPassword}}},
		}
		auth, _ := routerDependencyFlags(plans)
		if auth {
			t.Error("verify-password endpoint must not require auth")
		}
	})

	t.Run("PublishNeedsEventBus", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{TriggerKind: ir.TriggerSubscribe, Ops: []ir.Op{{Kind: ir.OpPublish}}},
		}
		auth, bus := routerDependencyFlags(plans)
		if auth || !bus {
			t.Errorf("got auth=%v bus=%v, want false,true", auth, bus)
		}
	})
}
