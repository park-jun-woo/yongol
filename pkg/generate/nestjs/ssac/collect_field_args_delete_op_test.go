//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_DeleteOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_DeleteOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpDelete,
		Delete: &ir.DeleteOp{
			Args: []ir.FieldArg{{Key: "id", Location: ir.LocPath}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}
