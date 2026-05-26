//ff:func feature=validate type=test control=sequence dimension=1 topic=ssac-structural
//ff:what s36MarkStaleAfterMutation — @put/@delete 후 stale 표시 검증 (매칭 모델 stale, 비 CRUD 스킵)

package ssac

import (
	"testing"

	parsessac "github.com/park-jun-woo/yongol/pkg/parser/ssac"
)

func TestS36MarkStaleAfterMutation(t *testing.T) {
	t.Run("Put", func(t *testing.T) {
		seq := parsessac.Sequence{Type: "put", Model: "Order.Update"}
		varType := map[string]string{"order": "Order", "user": "User"}
		stale := map[string]bool{}
		s36MarkStaleAfterMutation(seq, varType, stale)
		if !stale["order"] {
			t.Error("expected order to be stale after put")
		}
		if stale["user"] {
			t.Error("user should not be stale (different model)")
		}
	})
	t.Run("NonCRUDSkipped", func(t *testing.T) {
		seq := parsessac.Sequence{Type: "get", Model: "Order.FindByID"}
		varType := map[string]string{"order": "Order"}
		stale := map[string]bool{}
		s36MarkStaleAfterMutation(seq, varType, stale)
		if stale["order"] {
			t.Error("order should not be stale after get")
		}
	})
	t.Run("EmptyModelSkipped", func(t *testing.T) {
		seq := parsessac.Sequence{Type: "delete", Model: ""}
		varType := map[string]string{"order": "Order"}
		stale := map[string]bool{}
		s36MarkStaleAfterMutation(seq, varType, stale)
		if stale["order"] {
			t.Error("order should not be stale when model is empty")
		}
	})
}
