//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_PutOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_PutOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpPut,
		Put: &ir.PutOp{
			Args: []ir.FieldArg{{Key: "name", Location: ir.LocBody}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}
