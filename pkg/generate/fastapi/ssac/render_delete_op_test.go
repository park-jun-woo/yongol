//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderDeleteOp — DeleteOp → SQLAlchemy delete 문 렌더링

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderDeleteOp(t *testing.T) {
	t.Run("Nil", func(t *testing.T) {
		var b strings.Builder
		renderDeleteOp(&b, nil, "    ", "db")
		if b.String() != "" {
			t.Errorf("expected empty, got %q", b.String())
		}
	})
	t.Run("Basic", func(t *testing.T) {
		var b strings.Builder
		op := &ir.DeleteOp{
			Model: "work_item",
			Args:  []ir.FieldArg{{Location: ir.LocPath, ColumnName: "id"}},
		}
		renderDeleteOp(&b, op, "    ", "db")
		out := b.String()
		if !strings.Contains(out, "await db.execute(delete(WorkItem)") {
			t.Errorf("unexpected output: %q", out)
		}
		if !strings.HasPrefix(out, "    ") || !strings.HasSuffix(out, "\n") {
			t.Errorf("indent/newline missing: %q", out)
		}
	})
}
