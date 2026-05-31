//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferenceQuery_PaginationArgs
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

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
