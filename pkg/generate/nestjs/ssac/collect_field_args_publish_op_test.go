//ff:func feature=gen-nestjs type=test control=sequence
//ff:what TestCollectFieldArgs_PublishOp
package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestCollectFieldArgs_PublishOp(t *testing.T) {
	op := ir.Op{
		Kind: ir.OpPublish,
		Publish: &ir.PublishOp{
			Payload: []ir.FieldArg{{Key: "event_type"}},
			Options: []ir.FieldArg{{Key: "delay"}},
		},
	}
	got := collectFieldArgs(op)
	if len(got) != 2 {
		t.Fatalf("len = %d, want 2", len(got))
	}
}
