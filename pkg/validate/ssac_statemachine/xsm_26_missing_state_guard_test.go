//ff:func feature=validate type=test control=selection topic=states
//ff:what TestXsm26MissingStateGuard — XSM-26 state 전이 함수의 @state guard 누락 검증

package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXsm26MissingStateGuard(t *testing.T) {
	diagram := &statemachine.StateDiagram{
		ID:          "order",
		Transitions: []statemachine.Transition{{From: "draft", To: "active", Event: "CancelOrder"}},
	}

	t.Run("missing guard warns", func(t *testing.T) {
		fs := &yongol.Fullstack{
			StateDiagrams: []*statemachine.StateDiagram{diagram},
			ServiceFuncs: []ssac.ServiceFunc{{
				Name:      "CancelOrder",
				FileName:  "service/order/cancel.ssac",
				Sequences: []ssac.Sequence{{Type: "get"}},
			}},
		}
		if diags := xsm26MissingStateGuard(fs); len(diags) != 1 {
			t.Errorf("expected 1 diagnostic, got %d", len(diags))
		}
	})

	t.Run("guard present passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			StateDiagrams: []*statemachine.StateDiagram{diagram},
			ServiceFuncs: []ssac.ServiceFunc{{
				Name:      "CancelOrder",
				Sequences: []ssac.Sequence{{Type: "state", DiagramID: "order"}},
			}},
		}
		if diags := xsm26MissingStateGuard(fs); len(diags) != 0 {
			t.Errorf("expected no diagnostics, got %d", len(diags))
		}
	})
}
