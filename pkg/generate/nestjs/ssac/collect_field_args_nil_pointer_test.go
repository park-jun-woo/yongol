//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_NilPointer
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_NilPointer(t *testing.T) {
	// Nil inner pointer should not panic
	op := ir.Op{Kind: ir.OpGet, Get: nil}
	got := collectFieldArgs(op)
	if got != nil {
		t.Errorf("nil Get pointer should return nil, got %v", got)
	}
}
