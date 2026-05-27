package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestOpsReferenceQuery_True(t *testing.T) {
	ops := []ir.Op{
		{
			Kind: ir.OpGet,
			Get: &ir.GetOp{
				Args: []ir.FieldArg{
					{Key: "status", Location: ir.LocQuery},
				},
			},
		},
	}
	if !opsReferenceQuery(ops) {
		t.Error("should return true when LocQuery is present")
	}
}

func TestOpsReferenceQuery_False(t *testing.T) {
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
	if opsReferenceQuery(ops) {
		t.Error("should return false when no LocQuery")
	}
}

func TestOpsReferenceQuery_Empty(t *testing.T) {
	if opsReferenceQuery(nil) {
		t.Error("nil ops should return false")
	}
}

func TestOpsReferenceQuery_PaginationArgs(t *testing.T) {
	ops := []ir.Op{
		{
			Kind: ir.OpGet,
			Get: &ir.GetOp{
				PaginationArgs: []ir.FieldArg{
					{Key: "cursor", Location: ir.LocQuery},
				},
			},
		},
	}
	if !opsReferenceQuery(ops) {
		t.Error("should detect LocQuery in PaginationArgs")
	}
}
