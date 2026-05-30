//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestCollectPlanImports — Ops 슬라이스 순회하며 importData 누적

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectPlanImports(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		d := newImportData()
		collectPlanImports(&d, nil, "f")
		if len(d.Models) != 0 || d.HasAuth {
			t.Errorf("got %+v", d)
		}
	})
	t.Run("MultipleOps", func(t *testing.T) {
		d := newImportData()
		ops := []ir.Op{
			{Kind: ir.OpGet, Get: &ir.GetOp{Model: "User"}},
			{Kind: ir.OpPost, Post: &ir.PostOp{Model: "Item"}},
		}
		collectPlanImports(&d, ops, "f")
		if !d.UsesSelect || !d.Models["User"] || !d.Models["Item"] {
			t.Errorf("got %+v", d)
		}
	})
}
