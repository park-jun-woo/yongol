//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestXsm23TransitionToFunc — XSM-23 transition event → SSaC 함수 매칭 검증
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/ssac"
	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXsm23TransitionToFunc(t *testing.T) {
	diagram := &statemachine.StateDiagram{
		ID:          "order",
		Transitions: []statemachine.Transition{{From: "draft", To: "active", Event: "CancelOrder"}},
	}

	t.Run("matching func passes", func(t *testing.T) {
		fs := &yongol.Fullstack{
			StateDiagrams: []*statemachine.StateDiagram{diagram},
			ServiceFuncs:  []ssac.ServiceFunc{{Name: "CancelOrder"}},
		}
		if diags := xsm23TransitionToFunc(fs); len(diags) != 0 {
			t.Errorf("expected no diagnostics, got %d", len(diags))
		}
	})

	t.Run("unmatched event fires", func(t *testing.T) {
		fs := &yongol.Fullstack{
			StateDiagrams: []*statemachine.StateDiagram{diagram},
			ServiceFuncs:  nil,
		}
		if diags := xsm23TransitionToFunc(fs); len(diags) != 1 {
			t.Errorf("expected 1 diagnostic, got %d", len(diags))
		}
	})
}
