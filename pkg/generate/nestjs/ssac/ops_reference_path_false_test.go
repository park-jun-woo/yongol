//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestOpsReferencePath_False
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

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
