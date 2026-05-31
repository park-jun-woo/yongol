//ff:func feature=validate type=test control=sequence topic=states
//ff:what TestIsStatefulResource — isStatefulResource stateful 리소스 판정 분기 검증
package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestIsStatefulResource(t *testing.T) {
	ground := func() *rule.Ground {
		return &rule.Ground{Types: map[string]string{"DDL.default.value.workflows.status": "draft"}}
	}
	diagram := &statemachine.StateDiagram{ID: "workflow", InitialState: "draft"}

	t.Run("empty resource returns nil", func(t *testing.T) {
		if got := isStatefulResource("/{id}", []*statemachine.StateDiagram{diagram}, ground()); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})

	t.Run("no matching diagram returns nil", func(t *testing.T) {
		if got := isStatefulResource("/orders/{id}", []*statemachine.StateDiagram{diagram}, ground()); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})

	t.Run("matching stateful resource returns target", func(t *testing.T) {
		got := isStatefulResource("/workflows/{id}", []*statemachine.StateDiagram{diagram}, ground())
		if got == nil {
			t.Fatal("expected non-nil target")
		}
		if got.Table != "workflows" || got.StateColumn != "status" {
			t.Errorf("unexpected target: %+v", got)
		}
	})
}
