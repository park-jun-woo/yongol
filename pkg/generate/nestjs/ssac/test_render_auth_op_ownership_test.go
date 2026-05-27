//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestRenderAuthOpOwnership — Ownership DB lookup + authz.check 렌더링 검증

package ssac

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestRenderAuthOpOwnership(t *testing.T) {
	t.Run("WithOwnership", func(t *testing.T) {
		op := &ir.AuthOp{
			Action:   "ActivateWorkflow",
			Resource: "workflow",
			Inputs: []ir.FieldArg{
				{Key: "ResourceID", Location: ir.LocPath, ColumnName: "id", Source: "request", Field: ".ID"},
			},
			Ownership: &ir.OwnershipInfo{
				Table:       "workflows",
				OwnerColumn: "org_id",
				ResourcePK:  "id",
			},
		}
		var b strings.Builder
		renderAuthOp(&b, op, "      ")
		got := b.String()
		if !strings.Contains(got, "const owner = await tx.workflows.findUnique") {
			t.Errorf("expected ownership lookup, got: %s", got)
		}
		if !strings.Contains(got, "select: { org_id: true }") {
			t.Errorf("expected select owner column, got: %s", got)
		}
		if !strings.Contains(got, "resourceId: String(params.id)") {
			t.Errorf("expected resourceId, got: %s", got)
		}
		if !strings.Contains(got, "owners: { workflows: { org_id: owner?.org_id } }") {
			t.Errorf("expected owners map, got: %s", got)
		}
	})

	t.Run("WithoutOwnership", func(t *testing.T) {
		op := &ir.AuthOp{
			Action:   "ListWorkflows",
			Resource: "workflow",
			Inputs: []ir.FieldArg{
				{Key: "org_id", Location: ir.LocUser, ColumnName: "org_id", Source: "currentUser"},
			},
		}
		var b strings.Builder
		renderAuthOp(&b, op, "    ")
		got := b.String()
		if strings.Contains(got, "const owner") {
			t.Errorf("should not have ownership lookup, got: %s", got)
		}
		if !strings.Contains(got, "await this.authz.check") {
			t.Errorf("expected authz check, got: %s", got)
		}
	})

	t.Run("NilOp", func(t *testing.T) {
		var b strings.Builder
		renderAuthOp(&b, nil, "    ")
		if b.Len() != 0 {
			t.Error("expected empty for nil op")
		}
	})
}
