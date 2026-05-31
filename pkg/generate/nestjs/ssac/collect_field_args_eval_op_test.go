//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_EvalOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_EvalOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpEval,
		Eval: &ir.EvalOp{
			Args: []ir.FieldArg{{Key: "x", Location: ir.LocLiteral}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}
