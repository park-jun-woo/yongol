//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderPublishOp — PublishOp → event_bus.publish await 문 렌더링

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderPublishOp(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var b strings.Builder
		renderPublishOp(&b, nil, "    ")
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("NoPayload", func(t *testing.T) {
		var b strings.Builder
		renderPublishOp(&b, &ir.PublishOp{Topic: "order.completed"}, "    ")
		want := "    await event_bus.publish(\"order.completed\", {\n    })\n"
		if b.String() != want {
			t.Errorf("got %q, want %q", b.String(), want)
		}
	})
	t.Run("WithPayload", func(t *testing.T) {
		var b strings.Builder
		op := &ir.PublishOp{
			Topic: "order.completed",
			Payload: []ir.FieldArg{
				{Key: "order_id", Location: ir.LocVar, Source: "order", SourceColumn: "id"},
			},
		}
		renderPublishOp(&b, op, "    ")
		out := b.String()
		if !strings.Contains(out, `"order_id": order.id,`) {
			t.Errorf("missing payload entry: %q", out)
		}
	})
}
