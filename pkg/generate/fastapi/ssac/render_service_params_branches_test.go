//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderServiceParamsBranches — renderServiceParams 미커버 분기(subscribe+publish/bodyFallback/HTTP publish)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderServiceParamsBranches(t *testing.T) {
	t.Run("SubscribeWithPublish", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "on_event",
			TriggerKind: ir.TriggerSubscribe,
			Ops:         []ir.Op{{Kind: ir.OpPublish, Publish: &ir.PublishOp{Topic: "x"}}},
		}
		got := renderServiceParams(plan)
		want := "session: AsyncSession, payload: dict, event_bus: EventBus | None = None"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("BodyFallbackFromOps", func(t *testing.T) {
		// POST with LocBody op but no BodyFields metadata -> body: dict.
		plan := &ir.ServicePlan{
			OperationID: "create_thing",
			HTTPMethod:  "POST",
			TriggerKind: ir.TriggerHTTP,
			Ops: []ir.Op{
				{Kind: ir.OpPost, Post: &ir.PostOp{
					Model: "thing",
					Args:  []ir.FieldArg{{Location: ir.LocBody, ColumnName: "title"}},
				}},
			},
		}
		got := renderServiceParams(plan)
		if !strings.Contains(got, "body: dict") {
			t.Errorf("expected body: dict fallback, got %q", got)
		}
	})

	t.Run("HTTPWithPublishAddsEventBus", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "do_publish",
			HTTPMethod:  "POST",
			TriggerKind: ir.TriggerHTTP,
			Ops:         []ir.Op{{Kind: ir.OpPublish, Publish: &ir.PublishOp{Topic: "y"}}},
		}
		got := renderServiceParams(plan)
		if !strings.Contains(got, "event_bus: EventBus | None = None") {
			t.Errorf("expected event_bus param, got %q", got)
		}
	})
}
