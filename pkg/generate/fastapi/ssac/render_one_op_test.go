//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderOneOp — OpKind 분기 디스패치 (nil sub-pointer로 모든 case 도달)
package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderOneOp(t *testing.T) {
	kinds := []ir.OpKind{
		ir.OpGet, ir.OpPost, ir.OpPut, ir.OpDelete,
		ir.OpEmpty, ir.OpExists, ir.OpAuth, ir.OpState,
		ir.OpCall, ir.OpEval, ir.OpPublish, ir.OpVerifyPassword,
		ir.OpResponse,
	}
	for _, k := range kinds {
		t.Run("dispatch", func(t *testing.T) {
			var b strings.Builder
			// nil sub-pointers cause each renderer to early-return,
			// exercising the dispatch case without producing output.
			renderOneOp(&b, ir.Op{Kind: k}, "    ", "db")
			if b.String() != "" {
				t.Errorf("kind %d: expected empty for nil sub-op, got %q", k, b.String())
			}
		})
	}

	t.Run("UnknownKindNoop", func(t *testing.T) {
		var b strings.Builder
		renderOneOp(&b, ir.Op{Kind: ir.OpKind(9999)}, "    ", "db")
		if b.String() != "" {
			t.Errorf("expected empty for unknown kind, got %q", b.String())
		}
	})

	t.Run("RealDelete", func(t *testing.T) {
		var b strings.Builder
		op := ir.Op{Kind: ir.OpDelete, Delete: &ir.DeleteOp{
			Model: "item",
			Args:  []ir.FieldArg{{Location: ir.LocPath, ColumnName: "id"}},
		}}
		renderOneOp(&b, op, "    ", "db")
		if !strings.Contains(b.String(), "delete(Item)") {
			t.Errorf("expected delete render, got %q", b.String())
		}
	})
}
