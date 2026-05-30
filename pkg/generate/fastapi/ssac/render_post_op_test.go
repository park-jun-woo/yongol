//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderPostOp — PostOp → SQLAlchemy add/flush 문 렌더링

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderPostOp(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var b strings.Builder
		renderPostOp(&b, nil, "    ", "db")
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("Basic", func(t *testing.T) {
		var b strings.Builder
		op := &ir.PostOp{
			VarName: "item",
			Model:   "work_item",
			Args:    []ir.FieldArg{{Location: ir.LocBody, ColumnName: "title"}},
		}
		renderPostOp(&b, op, "    ", "db")
		out := b.String()
		for _, want := range []string{"item = WorkItem(", "db.add(item)", "await db.flush()"} {
			if !strings.Contains(out, want) {
				t.Errorf("missing %q in %q", want, out)
			}
		}
	})
}
