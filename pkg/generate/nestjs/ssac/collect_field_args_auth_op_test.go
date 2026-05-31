//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_AuthOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_AuthOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpAuth,
		Auth: &ir.AuthOp{
			Inputs: []ir.FieldArg{{Key: "org_id", Location: ir.LocPath}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
}
