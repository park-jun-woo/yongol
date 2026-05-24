//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what hasEmptyGuardAfter — empty/exists 가드 검출 (found/not found/다른 var) 검증

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestHasEmptyGuardAfter(t *testing.T) {
	t.Run("empty sequence returns false", func(t *testing.T) {
		if hasEmptyGuardAfter(nil, "x") {
			t.Error("expected false")
		}
	})

	t.Run("empty guard found", func(t *testing.T) {
		seqs := []parsessac.Sequence{{Type: "empty", Target: "user"}}
		if !hasEmptyGuardAfter(seqs, "user") {
			t.Error("expected true for @empty")
		}
	})

	t.Run("exists guard found", func(t *testing.T) {
		seqs := []parsessac.Sequence{{Type: "exists", Target: "user"}}
		if !hasEmptyGuardAfter(seqs, "user") {
			t.Error("expected true for @exists")
		}
	})

	t.Run("different var not found", func(t *testing.T) {
		seqs := []parsessac.Sequence{{Type: "empty", Target: "order"}}
		if hasEmptyGuardAfter(seqs, "user") {
			t.Error("expected false for different var")
		}
	})

	t.Run("non-guard type not found", func(t *testing.T) {
		seqs := []parsessac.Sequence{{Type: "get", Target: "user"}}
		if hasEmptyGuardAfter(seqs, "user") {
			t.Error("expected false for non-guard type")
		}
	})
}
