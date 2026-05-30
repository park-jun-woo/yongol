//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderService — ServicePlan → FastAPI service 함수 소스 생성

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderService(t *testing.T) {
	t.Run("NilPlan", func(t *testing.T) {
		_, err := RenderService(nil, nil)
		if err == nil {
			t.Fatal("expected error for nil plan")
		}
	})
	t.Run("Basic", func(t *testing.T) {
		plan := &ir.ServicePlan{
			OperationID: "delete_item",
			TriggerKind: ir.TriggerHTTP,
			Ops: []ir.Op{
				{Kind: ir.OpDelete, Delete: &ir.DeleteOp{
					Model: "item",
					Args:  []ir.FieldArg{{Location: ir.LocPath, ColumnName: "id"}},
				}},
			},
		}
		out, err := RenderService(plan, nil)
		if err != nil {
			t.Fatalf("RenderService: %v", err)
		}
		if !strings.Contains(out, "async def delete_item(") {
			t.Errorf("missing func def: %q", out)
		}
		if !strings.Contains(out, "delete(Item)") {
			t.Errorf("missing delete body: %q", out)
		}
	})
}
