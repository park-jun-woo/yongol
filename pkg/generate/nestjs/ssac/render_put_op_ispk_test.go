//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderPutOpIsPK — IsPK 기반 where/data 분리 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderPutOpIsPK(t *testing.T) {
	t.Run("PKSplitsCorrectly", func(t *testing.T) {
		op := &ir.PutOp{
			Model: "Workflow",
			Args: []ir.FieldArg{
				{Location: ir.LocPath, ColumnName: "id", IsPK: true, Source: "request", Field: ".ID"},
				{Location: ir.LocBody, ColumnName: "title", Source: "request", Field: ".Title"},
				{Location: ir.LocBody, ColumnName: "status", Source: "request", Field: ".Status"},
			},
		}
		var b strings.Builder
		renderPutOp(&b, op, "    ", "tx")
		got := b.String()
		if !strings.Contains(got, "where: { id: params.id }") {
			t.Errorf("expected where with PK, got: %s", got)
		}
		if !strings.Contains(got, "data: { title: body.title, status: body.status }") {
			t.Errorf("expected data without PK, got: %s", got)
		}
	})

	t.Run("NilOp", func(t *testing.T) {
		var b strings.Builder
		renderPutOp(&b, nil, "    ", "tx")
		if b.Len() != 0 {
			t.Error("expected empty output for nil op")
		}
	})

	t.Run("NoPKFallback", func(t *testing.T) {
		op := &ir.PutOp{
			Model: "Workflow",
			Args: []ir.FieldArg{
				{Location: ir.LocBody, ColumnName: "title", Source: "request", Field: ".Title"},
			},
		}
		var b strings.Builder
		renderPutOp(&b, op, "    ", "tx")
		got := b.String()
		// No PK arg → fallback where
		if !strings.Contains(got, "where: { id: params.id }") {
			t.Errorf("expected default PK fallback, got: %s", got)
		}
	})
}
