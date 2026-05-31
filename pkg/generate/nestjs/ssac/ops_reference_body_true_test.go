//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferenceBody_True
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
