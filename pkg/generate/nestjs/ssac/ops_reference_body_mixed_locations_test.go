//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferenceBody_MixedLocations
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

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
