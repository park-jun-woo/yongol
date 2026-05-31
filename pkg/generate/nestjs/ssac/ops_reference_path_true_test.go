//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferencePath_True
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
