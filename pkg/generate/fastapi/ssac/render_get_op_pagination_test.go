//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderGetOpPagination — PaginationArgs → SQLAlchemy limit/offset 렌더링 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderGetOpPagination(t *testing.T) {
	t.Run("WithPaginationArgs", func(t *testing.T) {
		op := &ir.GetOp{
			VarName: "workflows",
			VarType: "Workflow",
			Model:   "workflow",
			IsList:  true,
			Args: []ir.FieldArg{
				{Location: ir.LocUser, ColumnName: "org_id", Source: "currentUser", Field: ".OrgID"},
			},
			PaginationArgs: []ir.FieldArg{
				{Location: ir.LocQuery, ColumnName: "per_page", Source: "request", Field: ".PerPage"},
				{Location: ir.LocQuery, ColumnName: "offset", Source: "request", Field: ".Offset"},
			},
		}
		var b strings.Builder
		renderGetOp(&b, op, "    ", "session")
		got := b.String()
		if !strings.Contains(got, "scalars().all()") {
			t.Errorf("expected scalars().all() for IsList, got: %s", got)
		}
		if !strings.Contains(got, `Workflow.org_id == current_user["org_id"]`) {
			t.Errorf("expected where clause, got: %s", got)
		}
		if !strings.Contains(got, ".limit(per_page)") {
			t.Errorf("expected limit from per_page, got: %s", got)
		}
		if !strings.Contains(got, ".offset(offset)") {
			t.Errorf("expected offset, got: %s", got)
		}
	})

	t.Run("FindUniqueNoVarShadow", func(t *testing.T) {
		op := &ir.GetOp{
			VarName: "wf",
			VarType: "Workflow",
			Model:   "workflow",
			IsList:  false,
			Args: []ir.FieldArg{
				{Location: ir.LocPath, ColumnName: "id", Source: "request", Field: ".ID"},
			},
		}
		var b strings.Builder
		renderGetOp(&b, op, "    ", "session")
		got := b.String()
		if !strings.Contains(got, "scalars().first()") {
			t.Errorf("expected scalars().first(), got: %s", got)
		}
		if !strings.Contains(got, "wf = result.scalars()") {
			t.Errorf("expected VarName used directly, got: %s", got)
		}
	})

	t.Run("NilOp", func(t *testing.T) {
		var b strings.Builder
		renderGetOp(&b, nil, "    ", "session")
		if b.Len() != 0 {
			t.Error("expected empty for nil op")
		}
	})
}
