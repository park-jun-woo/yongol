//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderStateOpAllowed — AllowedFromStates 기반 전이 맵 렌더링 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderStateOpAllowed(t *testing.T) {
	t.Run("WithAllowedStates", func(t *testing.T) {
		op := &ir.StateOp{
			Diagram:    "workflows",
			Transition: "ActivateWorkflow",
			Inputs: []ir.FieldArg{
				{Key: "status", Location: ir.LocVar, Source: "wf", ColumnName: "status"},
			},
			Message:           "Cannot activate workflow",
			StatusCode:        409,
			AllowedFromStates: []string{"draft", "paused"},
		}
		var b strings.Builder
		renderStateOp(&b, op, "    ")
		got := b.String()
		if !strings.Contains(got, `"draft": True`) {
			t.Errorf("expected draft in allowed map, got: %s", got)
		}
		if !strings.Contains(got, `"paused": True`) {
			t.Errorf("expected paused in allowed map, got: %s", got)
		}
		if strings.Contains(got, "TODO") {
			t.Errorf("should not have TODO when AllowedFromStates populated, got: %s", got)
		}
		if !strings.Contains(got, "wf.status not in allowed_activate_workflow") {
			t.Errorf("expected status check, got: %s", got)
		}
	})

	t.Run("WithoutAllowedStates", func(t *testing.T) {
		op := &ir.StateOp{
			Diagram:    "workflows",
			Transition: "ActivateWorkflow",
			Inputs: []ir.FieldArg{
				{Key: "status", Location: ir.LocVar, Source: "wf", ColumnName: "status"},
			},
			Message:    "Cannot activate",
			StatusCode: 409,
		}
		var b strings.Builder
		renderStateOp(&b, op, "    ")
		got := b.String()
		if !strings.Contains(got, "TODO") {
			t.Errorf("expected TODO when no AllowedFromStates, got: %s", got)
		}
	})

	t.Run("NilOp", func(t *testing.T) {
		var b strings.Builder
		renderStateOp(&b, nil, "    ")
		if b.Len() != 0 {
			t.Error("expected empty for nil op")
		}
	})
}
