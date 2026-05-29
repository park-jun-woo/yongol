//ff:func feature=validate-contract type=test control=sequence topic=preserve-safety
//ff:what TestMarkLenGuard — CallExpr 가 len(x) 이면 x 를 guarded 집합에 추가 검증

package contract

import "testing"

func TestMarkLenGuard(t *testing.T) {
	t.Run("len of ident", func(t *testing.T) {
		guarded := map[string]bool{}
		if !markLenGuard(mustCall(t, "len(xs)"), guarded) {
			t.Fatal("expected match for len(xs)")
		}
		if !guarded["xs"] {
			t.Fatalf("expected xs guarded, got %v", guarded)
		}
	})
	t.Run("non len call", func(t *testing.T) {
		guarded := map[string]bool{}
		if markLenGuard(mustCall(t, "cap(xs)"), guarded) {
			t.Fatal("cap should not match")
		}
		if len(guarded) != 0 {
			t.Fatalf("expected empty guard, got %v", guarded)
		}
	})
	t.Run("len of index expr ignored", func(t *testing.T) {
		guarded := map[string]bool{}
		if markLenGuard(mustCall(t, "len(m[k])"), guarded) {
			t.Fatal("len(m[k]) arg is not a simple ident; should be ignored")
		}
	})
	t.Run("nil call", func(t *testing.T) {
		if markLenGuard(nil, map[string]bool{}) {
			t.Fatal("nil call should return false")
		}
	})
}
