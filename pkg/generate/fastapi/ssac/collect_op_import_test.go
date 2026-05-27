//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestCollectOpImport — OpAuth ownership 모델 등록 + UsesSelect 플래그 검증

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectOpImport(t *testing.T) {
	t.Run("AuthWithOwnership", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpAuth,
			Auth: &ir.AuthOp{
				Action:   "ActivateWorkflow",
				Resource: "workflow",
				Ownership: &ir.OwnershipInfo{
					Table:       "workflows",
					OwnerColumn: "org_id",
					ResourcePK:  "id",
				},
			},
		}
		collectOpImport(&d, op, "workflow")
		if !d.HasAuth {
			t.Error("expected HasAuth true")
		}
		if !d.UsesSelect {
			t.Error("expected UsesSelect true for ownership lookup")
		}
		if !d.Models["Workflow"] {
			t.Errorf("expected singular model Workflow in Models, got %v", d.Models)
		}
	})

	t.Run("AuthWithoutOwnership", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpAuth,
			Auth: &ir.AuthOp{
				Action:   "ListWorkflows",
				Resource: "workflow",
			},
		}
		collectOpImport(&d, op, "workflow")
		if !d.HasAuth {
			t.Error("expected HasAuth true")
		}
		if d.UsesSelect {
			t.Error("expected UsesSelect false without ownership")
		}
		if len(d.Models) != 0 {
			t.Errorf("expected empty Models without ownership, got %v", d.Models)
		}
	})

	t.Run("AuthNilOp", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpAuth,
		}
		collectOpImport(&d, op, "workflow")
		if !d.HasAuth {
			t.Error("expected HasAuth true even with nil Auth")
		}
		if d.UsesSelect {
			t.Error("expected UsesSelect false with nil Auth")
		}
	})

	t.Run("OwnershipSingularizesIES", func(t *testing.T) {
		d := importData{
			Models:  make(map[string]bool),
			ExtPkgs: make(map[string]map[string]bool),
		}
		op := ir.Op{
			Kind: ir.OpAuth,
			Auth: &ir.AuthOp{
				Action:   "UpdateCategory",
				Resource: "category",
				Ownership: &ir.OwnershipInfo{
					Table:       "categories",
					OwnerColumn: "owner_id",
					ResourcePK:  "id",
				},
			},
		}
		collectOpImport(&d, op, "category")
		if !d.Models["Category"] {
			t.Errorf("expected singular model Category from categories, got %v", d.Models)
		}
	})
}
