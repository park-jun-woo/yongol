//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestHasVerifyPasswordOp — Op 배열 내 VerifyPassword 연산 존재 여부

package ssac

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/generate/ir"
)

func TestHasVerifyPasswordOp(t *testing.T) {
	t.Run("Found", func(t *testing.T) {
		ops := []ir.Op{{Kind: ir.OpGet}, {Kind: ir.OpVerifyPassword}}
		if !hasVerifyPasswordOp(ops) {
			t.Error("expected true")
		}
	})
	t.Run("NotFound", func(t *testing.T) {
		ops := []ir.Op{{Kind: ir.OpGet}, {Kind: ir.OpPost}}
		if hasVerifyPasswordOp(ops) {
			t.Error("expected false")
		}
	})
	t.Run("Empty", func(t *testing.T) {
		if hasVerifyPasswordOp(nil) {
			t.Error("expected false for empty")
		}
	})
}
