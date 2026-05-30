//ff:func feature=validate type=test control=selection topic=states
//ff:what TestXms25StateEvent — XMS-25 @state transition → diagram event 검증

package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXms25StateEvent(t *testing.T) {
	t.Run("valid transition passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			StateDiagrams: []*statemachine.StateDiagram{{
				ID:          "order",
				Transitions: []statemachine.Transition{{From: "draft", To: "active", Event: "cancel"}},
			}},
			ServiceFuncs: []ssac.ServiceFunc{{
				Name:      "CancelOrder",
				Sequences: []ssac.Sequence{{Type: "state", DiagramID: "order", Transition: "cancel"}},
			}},
		}
		if diags := xms25StateEvent(fs); len(diags) != 0 {
			t.Errorf("expected no diagnostics, got %d: %+v", len(diags), diags)
		}
	})

	t.Run("invalid transition fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			StateDiagrams: []*statemachine.StateDiagram{{
				ID:          "order",
				Transitions: []statemachine.Transition{{From: "draft", To: "active", Event: "cancel"}},
			}},
			ServiceFuncs: []ssac.ServiceFunc{{
				Name:      "ShipOrder",
				FileName:  "service/order/ship.ssac",
				Sequences: []ssac.Sequence{{Type: "state", DiagramID: "order", Transition: "ship", Line: 3}},
			}},
		}
		if diags := xms25StateEvent(fs); len(diags) == 0 {
			t.Error("expected diagnostics for invalid transition")
		}
	})
}
