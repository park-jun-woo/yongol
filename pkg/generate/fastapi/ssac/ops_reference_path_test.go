//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestOpsReference — opsReferenceBody/Path/Query FieldArg Location 검사 검증
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestOpsReferencePath(t *testing.T) {
	t.Run("PathArgPresent", func(t *testing.T) {
		ops := []ir.Op{
			{Kind: ir.OpGet, Get: &ir.GetOp{
				Args: []ir.FieldArg{{Location: ir.LocPath, Key: "id"}},
			}},
		}
		if !opsReferencePath(ops) {
			t.Error("expected true for LocPath arg")
		}
	})

	t.Run("NoPathArg", func(t *testing.T) {
		ops := []ir.Op{
			{Kind: ir.OpPost, Post: &ir.PostOp{
				Args: []ir.FieldArg{{Location: ir.LocBody, Key: "title"}},
			}},
		}
		if opsReferencePath(ops) {
			t.Error("expected false when no LocPath arg")
		}
	})

	t.Run("NilOps", func(t *testing.T) {
		if opsReferencePath(nil) {
			t.Error("expected false for nil ops")
		}
	})
}
