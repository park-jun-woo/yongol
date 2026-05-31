//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_ExistsOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_ExistsOp(t *testing.T) {
	op := ir.Op{
		Kind:   ir.OpExists,
		Exists: &ir.ExistsOp{VarName: "existing"},
	}
	got := collectFieldArgs(op)
	if got != nil {
		t.Errorf("ExistsOp should return nil, got %v", got)
	}
}
