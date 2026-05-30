//ff:func feature=gen-fastapi type=test control=iteration dimension=1
//ff:what TestRenderSAWhere — FieldArg 배열 → SQLAlchemy .where() 절 문자열

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderSAWhere(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		if got := renderSAWhere("work_item", nil); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
	t.Run("SingleClause", func(t *testing.T) {
		args := []ir.FieldArg{{Location: ir.LocPath, ColumnName: "id"}}
		got := renderSAWhere("work_item", args)
		want := ".where(WorkItem.id == id)"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
	t.Run("MultipleClauses", func(t *testing.T) {
		args := []ir.FieldArg{
			{Location: ir.LocPath, ColumnName: "id"},
			{Location: ir.LocUser, ColumnName: "org_id"},
		}
		got := renderSAWhere("work_item", args)
		want := `.where(WorkItem.id == id, WorkItem.org_id == current_user["org_id"])`
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
