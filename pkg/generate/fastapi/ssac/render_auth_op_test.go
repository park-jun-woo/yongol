//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderAuthOpOwnership — Ownership DB lookup + authz_check 호출 렌더링 검증

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
		renderAuthOp(&b, op, "    ")
		got := b.String()
		if !strings.Contains(got, "owner_row = await session.execute") {
			t.Errorf("expected ownership lookup, got: %s", got)
		}
		if !strings.Contains(got, "Workflow.org_id") {
			t.Errorf("expected singular model select owner column, got: %s", got)
		}
		if !strings.Contains(got, "scalar_one_or_none") {
			t.Errorf("expected scalar_one_or_none, got: %s", got)
		}
		if !strings.Contains(got, `owners={"workflow": {"org_id": owner}}`) {
			t.Errorf("expected owners map with resource key, got: %s", got)
		}
		if !strings.Contains(got, "resource_id=str(id)") {
			t.Errorf("expected resource_id, got: %s", got)
		}
		// Phase024-3a: ResourceID must NOT appear as a duplicate keyword
		count := strings.Count(got, "resource_id=")
		if count != 1 {
			t.Errorf("expected exactly 1 resource_id=, got %d in: %s", count, got)
		}
	})

	t.Run("WithoutOwnershipButResourceID", func(t *testing.T) {
		op := &ir.AuthOp{
			Action:   "GetWorkflow",
			Resource: "workflow",
			Inputs: []ir.FieldArg{
				{Key: "ResourceID", Location: ir.LocPath, ColumnName: "id", Source: "request", Field: ".ID"},
			},
		}
		var b strings.Builder
		renderAuthOp(&b, op, "    ")
		got := b.String()
		if !strings.Contains(got, "resource_id=str(id)") {
			t.Errorf("expected fallback resource_id, got: %s", got)
		}
		count := strings.Count(got, "resource_id=")
		if count != 1 {
			t.Errorf("expected exactly 1 resource_id=, got %d in: %s", count, got)
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
		if strings.Contains(got, "owner_row") {
			t.Errorf("should not have ownership lookup, got: %s", got)
		}
		if !strings.Contains(got, "await authz_check") {
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
