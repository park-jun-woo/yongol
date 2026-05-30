//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestWriteServiceFuncTransaction — UsesTransaction 시 async with session.begin() 블록 렌더링

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestWriteServiceFuncTransaction(t *testing.T) {
	plan := &ir.ServicePlan{
		OperationID:     "delete_item",
		TriggerKind:     ir.TriggerHTTP,
		HTTPMethod:      "DELETE",
		UsesTransaction: true,
		Ops: []ir.Op{
			{Kind: ir.OpDelete, Delete: &ir.DeleteOp{
				Model: "item",
				Args:  []ir.FieldArg{{Location: ir.LocPath, ColumnName: "id"}},
			}},
		},
	}
	var b strings.Builder
	writeServiceFunc(&b, plan)
	got := b.String()
	if !strings.Contains(got, "async with session.begin():") {
		t.Errorf("expected transaction block, got:\n%s", got)
	}
	// Ops should be rendered with the deeper (8-space) indent.
	if !strings.Contains(got, "        await session.execute(delete(Item)") {
		t.Errorf("expected indented op body, got:\n%s", got)
	}
}
