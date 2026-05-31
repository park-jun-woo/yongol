//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferenceQuery_True
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
