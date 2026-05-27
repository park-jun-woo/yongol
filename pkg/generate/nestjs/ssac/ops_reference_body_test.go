package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestOpsReferenceBody_True(t *testing.T) {
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
	if !opsReferenceBody(ops) {
		t.Error("should return true when LocBody is present")
	}
}

func TestOpsReferenceBody_False(t *testing.T) {
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
	if opsReferenceBody(ops) {
		t.Error("should return false when no LocBody")
	}
}

func TestOpsReferenceBody_Empty(t *testing.T) {
	if opsReferenceBody(nil) {
		t.Error("nil ops should return false")
	}
}

func TestOpsReferenceBody_MixedLocations(t *testing.T) {
	ops := []ir.Op{
		{
			Kind: ir.OpGet,
			Get: &ir.GetOp{
				Args: []ir.FieldArg{
					{Key: "id", Location: ir.LocPath},
				},
			},
		},
		{
			Kind: ir.OpPut,
			Put: &ir.PutOp{
				Args: []ir.FieldArg{
					{Key: "name", Location: ir.LocBody},
				},
			},
		},
	}
	if !opsReferenceBody(ops) {
		t.Error("should return true when LocBody is in any op")
	}
}
