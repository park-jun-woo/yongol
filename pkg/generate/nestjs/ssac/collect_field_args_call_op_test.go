//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_CallOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_CallOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpCall,
		Call: &ir.CallOp{
			Args: []ir.FieldArg{{Key: "input", Location: ir.LocVar}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}
