//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferenceBody_False
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

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
