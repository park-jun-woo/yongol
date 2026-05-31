//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_StateOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_StateOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpState,
		State: &ir.StateOp{
			Inputs: []ir.FieldArg{{Key: "status", Location: ir.LocBody}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}
