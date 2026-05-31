//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_EmptyOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_EmptyOp(t *testing.T) {
	op := ir.Op{
		Kind:  ir.OpEmpty,
		Empty: &ir.EmptyOp{VarName: "item"},
	}
	got := collectFieldArgs(op)
	if got != nil {
		t.Errorf("EmptyOp should return nil, got %v", got)
	}
}
