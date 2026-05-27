package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestOpsReferencePath_True(t *testing.T) {
	ops := []ir.Op{
		{
			Kind: ir.OpGet,
			Get: &ir.GetOp{
				Args: []ir.FieldArg{
					{Key: "id", Location: ir.LocPath},
				},
			},
		},
	}
	if !opsReferencePath(ops) {
		t.Error("should return true when LocPath is present")
	}
}

func TestOpsReferencePath_False(t *testing.T) {
	ops := []ir.Op{
		{
			Kind: ir.OpPost,
			Post: &ir.PostOp{
				Args: []ir.FieldArg{
					{Key: "title", Location: ir.LocBody},
				},
			},
		},
	}
	if opsReferencePath(ops) {
		t.Error("should return false when no LocPath")
	}
}

func TestOpsReferencePath_Empty(t *testing.T) {
	if opsReferencePath(nil) {
		t.Error("nil ops should return false")
	}
}
