//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestHasPublishOp — Op 배열 내 Publish 연산 존재 여부

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasPublishOp(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		ops := []ir.Op{{Kind: ir.OpGet}, {Kind: ir.OpPublish}}
		if !hasPublishOp(ops) {
			t.Error("expected true")
		}
	})
	t.Run("NotFound", func(t *testing.T) {
		ops := []ir.Op{{Kind: ir.OpGet}, {Kind: ir.OpPost}}
		if hasPublishOp(ops) {
			t.Error("expected false")
		}
	})
	t.Run("Empty", func(t *testing.T) {
		if hasPublishOp(nil) {
			t.Error("expected false for empty")
		}
	})
}
