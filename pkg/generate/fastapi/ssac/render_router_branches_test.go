//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderRouterBranches — RenderRouter 미커버 분기(빈 feature/schema import/subscribe)

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderRouterBranches(t *testing.T) {
	t.Run("EmptyFeatureError", func(t *testing.T) {
		_, err := RenderRouter("", nil)
		if err == nil {
			t.Fatal("expected error for empty feature")
		}
	})

	t.Run("SchemaImportForBody", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "create_item",
				HTTPMethod:  "POST",
				TriggerKind: ir.TriggerHTTP,
				URLPath:     "/item",
				Feature:     "item",
				BodyFields:  []ir.BodyFieldMeta{{}},
				Ops:         []ir.Op{{Kind: ir.OpVerifyPassword, VerifyPW: &ir.VerifyPasswordOp{Model: "User"}}},
			},
		}
		got, err := RenderRouter("item", plans)
		if err != nil {
			t.Fatalf("RenderRouter: %v", err)
		}
		if !strings.Contains(got, "from app.schemas.item import CreateItemRequest") {
			t.Errorf("expected schema import, got:\n%s", got)
		}
		// verify-password endpoint should not pull in get_current_user.
		if strings.Contains(got, "get_current_user") {
			t.Errorf("verify-password only plan should skip auth import, got:\n%s", got)
		}
	})

	t.Run("SubscribeHandler", func(t *testing.T) {
		plans := []*ir.ServicePlan{
			{
				OperationID: "on_order_paid",
				TriggerKind: ir.TriggerSubscribe,
				Feature:     "billing",
				Topic:       "order.paid",
				Ops:         []ir.Op{{Kind: ir.OpGet, Get: &ir.GetOp{Model: "Order"}}},
			},
		}
		got, err := RenderRouter("billing", plans)
		if err != nil {
			t.Fatalf("RenderRouter: %v", err)
		}
		if !strings.Contains(got, "# Subscribe handler for topic: order.paid") {
			t.Errorf("expected subscribe handler, got:\n%s", got)
		}
		if strings.Contains(got, "get_current_user") {
			t.Errorf("subscribe-only router should not import auth, got:\n%s", got)
		}
	})
}
