//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderGetOpPagination — PaginationArgs → Prisma take/skip/cursor 렌더링 검증

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
			Model:   "Workflow",
			IsList:  true,
			Args: []ir.FieldArg{
				{Location: ir.LocUser, ColumnName: "org_id", Source: "currentUser", Field: ".OrgID"},
			},
			PaginationArgs: []ir.FieldArg{
				{Location: ir.LocQuery, ColumnName: "per_page", Source: "request", Field: ".PerPage"},
				{Location: ir.LocQuery, ColumnName: "cursor", Source: "request", Field: ".Cursor"},
			},
		}
		var b strings.Builder
		renderGetOp(&b, op, "    ", "this.prisma")
		got := b.String()
		if !strings.Contains(got, "findMany") {
			t.Errorf("expected findMany for IsList, got: %s", got)
		}
		if !strings.Contains(got, "where: { org_id: user.org_id }") {
			t.Errorf("expected where clause, got: %s", got)
		}
		if !strings.Contains(got, "take: query.per_page") {
			t.Errorf("expected take from per_page, got: %s", got)
		}
		if !strings.Contains(got, "cursor: query.cursor") {
			t.Errorf("expected cursor, got: %s", got)
		}
	})

	t.Run("FindUniqueNoVarShadow", func(t *testing.T) {
		op := &ir.GetOp{
			VarName: "wf",
			VarType: "Workflow",
			Model:   "Workflow",
			IsList:  false,
			Args: []ir.FieldArg{
				{Location: ir.LocPath, ColumnName: "id", Source: "request", Field: ".ID"},
			},
		}
		var b strings.Builder
		renderGetOp(&b, op, "    ", "this.prisma")
		got := b.String()
		if !strings.Contains(got, "findUnique") {
			t.Errorf("expected findUnique, got: %s", got)
		}
		if !strings.Contains(got, "const wf =") {
			t.Errorf("expected VarName used directly, got: %s", got)
		}
	})

	t.Run("CountQuery", func(t *testing.T) {
		op := &ir.GetOp{
			VarName: "total",
			VarType: "int64",
			Model:   "AuditLog",
			IsCount: true,
			Args: []ir.FieldArg{
				{Location: ir.LocUser, ColumnName: "org_id", Source: "currentUser", Field: ".OrgID"},
			},
		}
		var b strings.Builder
		renderGetOp(&b, op, "    ", "this.prisma")
		got := b.String()
		if !strings.Contains(got, ".count(") {
			t.Errorf("expected count() for IsCount, got: %s", got)
		}
		if strings.Contains(got, "findUnique") || strings.Contains(got, "findMany") {
			t.Errorf("should not use findUnique/findMany for count, got: %s", got)
		}
		if !strings.Contains(got, "where: { org_id: user.org_id }") {
			t.Errorf("expected where clause in count, got: %s", got)
		}
	})

	t.Run("CountQueryNoArgs", func(t *testing.T) {
		op := &ir.GetOp{
			VarName: "total",
			VarType: "int64",
			Model:   "Item",
			IsCount: true,
		}
		var b strings.Builder
		renderGetOp(&b, op, "    ", "this.prisma")
		got := b.String()
		if !strings.Contains(got, ".count()") {
			t.Errorf("expected count() with no args, got: %s", got)
		}
	})

	t.Run("NilOp", func(t *testing.T) {
		var b strings.Builder
		renderGetOp(&b, nil, "    ", "tx")
		if b.Len() != 0 {
			t.Error("expected empty for nil op")
		}
	})
}
