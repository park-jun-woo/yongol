//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_GetOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_GetOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpGet,
		Get: &ir.GetOp{
			Args:           []ir.FieldArg{{Key: "id", Location: ir.LocPath}},
			PaginationArgs: []ir.FieldArg{{Key: "cursor", Location: ir.LocQuery}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
	if got[0].Key != "id" {
		t.Errorf("got[0].Key = %q, want %q", got[0].Key, "id")
	}
	if got[1].Key != "cursor" {
		t.Errorf("got[1].Key = %q, want %q", got[1].Key, "cursor")
	}
}
