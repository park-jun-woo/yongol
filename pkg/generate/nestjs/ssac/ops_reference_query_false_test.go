//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferenceQuery_False
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

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
