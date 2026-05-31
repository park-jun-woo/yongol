//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_ResponseOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_ResponseOp(t *testing.T) {
	op := ir.Op{
		Kind:     ir.OpResponse,
		Response: &ir.ResponseOp{SingleVar: "result"},
	}
	got := collectFieldArgs(op)
	if got != nil {
		t.Errorf("ResponseOp should return nil, got %v", got)
	}
}
