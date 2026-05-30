//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderOpsBody — Op 배열 일괄 렌더링

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderOpsBody(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		var b strings.Builder
		renderOpsBody(&b, nil, "    ", "db")
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("Multiple", func(t *testing.T) {
		var b strings.Builder
		ops := []ir.Op{
			{Kind: ir.OpEmpty, Empty: &ir.EmptyOp{VarName: "x", StatusCode: 404, Message: "nf"}},
			{Kind: ir.OpDelete, Delete: &ir.DeleteOp{Model: "item", Args: []ir.FieldArg{{Location: ir.LocPath, ColumnName: "id"}}}},
		}
		renderOpsBody(&b, ops, "    ", "db")
		out := b.String()
		if !strings.Contains(out, "if not x:") || !strings.Contains(out, "delete(Item)") {
			t.Errorf("unexpected output: %q", out)
		}
	})
}
