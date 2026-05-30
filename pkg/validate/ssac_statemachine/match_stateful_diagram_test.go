//ff:func feature=validate type=test control=selection topic=states
//ff:what TestMatchStatefulDiagram — matchStatefulDiagram 단일 diagram stateful 대응 분기 검증

package ssac_statemachine

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/statemachine"
	"github.com/park-jun-woo/yongol/pkg/rule"
)

func TestMatchStatefulDiagram(t *testing.T) {
	groundOK := &rule.Ground{Types: map[string]string{"DDL.default.value.workflows.status": "draft"}}

	t.Run("nil diagram", func(t *testing.T) {
		if got := matchStatefulDiagram(nil, "workflows", "workflow", groundOK); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})

	t.Run("table mismatch", func(t *testing.T) {
		d := &statemachine.StateDiagram{ID: "order", InitialState: "draft"}
		if got := matchStatefulDiagram(d, "workflows", "workflow", groundOK); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})

	t.Run("empty initial state", func(t *testing.T) {
		d := &statemachine.StateDiagram{ID: "workflow", InitialState: ""}
		if got := matchStatefulDiagram(d, "workflows", "workflow", groundOK); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})

	t.Run("nil ground", func(t *testing.T) {
		d := &statemachine.StateDiagram{ID: "workflow", InitialState: "draft"}
		if got := matchStatefulDiagram(d, "workflows", "workflow", nil); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})

	t.Run("default mismatch", func(t *testing.T) {
		d := &statemachine.StateDiagram{ID: "workflow", InitialState: "active"}
		if got := matchStatefulDiagram(d, "workflows", "workflow", groundOK); got != nil {
			t.Errorf("expected nil, got %+v", got)
		}
	})

	t.Run("match", func(t *testing.T) {
		d := &statemachine.StateDiagram{ID: "workflow", InitialState: "draft"}
		got := matchStatefulDiagram(d, "workflows", "workflow", groundOK)
		if got == nil {
			t.Fatal("expected non-nil target")
		}
		if got.Table != "workflows" || got.StateColumn != "status" || got.Model != "Workflow" {
			t.Errorf("unexpected target: %+v", got)
		}
	})
}
