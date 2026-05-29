//ff:func feature=gen-fastapi type=test control=sequence
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
			Model: "workflow",
			Args: []ir.FieldArg{
				{Location: ir.LocPath, ColumnName: "id", IsPK: true, Source: "request", Field: ".ID"},
				{Location: ir.LocBody, ColumnName: "title", Source: "request", Field: ".Title"},
				{Location: ir.LocBody, ColumnName: "status", Source: "request", Field: ".Status"},
			},
		}
		var b strings.Builder
		renderPutOp(&b, op, "    ", "session")
		got := b.String()
		if !strings.Contains(got, "Workflow.id == id") {
			t.Errorf("expected where with PK path param, got: %s", got)
		}
		if !strings.Contains(got, "title=body.title") {
			t.Errorf("expected data with body fields, got: %s", got)
		}
		if !strings.Contains(got, "status=body.status") {
			t.Errorf("expected status in data, got: %s", got)
		}
	})

	t.Run("NilOp", func(t *testing.T) {
		var b strings.Builder
		renderPutOp(&b, nil, "    ", "session")
		if b.Len() != 0 {
			t.Error("expected empty output for nil op")
		}
	})

	t.Run("NoPKFallback", func(t *testing.T) {
		op := &ir.PutOp{
			Model: "workflow",
			Args: []ir.FieldArg{
				{Location: ir.LocBody, ColumnName: "title", Source: "request", Field: ".Title"},
			},
		}
		var b strings.Builder
		renderPutOp(&b, op, "    ", "session")
		got := b.String()
		// No PK arg → all go to data, no where clause
		if !strings.Contains(got, "title=body.title") {
			t.Errorf("expected data clause, got: %s", got)
		}
	})
}
