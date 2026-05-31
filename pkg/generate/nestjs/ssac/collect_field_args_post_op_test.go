//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_PostOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_PostOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpPost,
		Post: &ir.PostOp{
			Args: []ir.FieldArg{{Key: "title", Location: ir.LocBody}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 1 {
		t.Fatalf("len = %d, want 1", len(got))
	}
	if got[0].Key != "title" {
		t.Errorf("got[0].Key = %q, want %q", got[0].Key, "title")
	}
}
