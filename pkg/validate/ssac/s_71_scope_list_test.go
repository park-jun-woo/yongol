//ff:func feature=validate type=test control=sequence dimension=1 topic=ssac-structural
//ff:what S-71 — s71ScopeList 단위 테스트 (scope map → 정렬된 문자열)

package ssac

import "testing"

func TestS71ScopeList(t *testing.T) {
	t.Run("Sorted_output", func(t *testing.T) {
		got := s71ScopeList(map[string]bool{"request": true, "course": true, "currentUser": true})
		want := "course, currentUser, request"
		if got != want {
			t.Errorf("s71ScopeList() = %q, want %q", got, want)
		}
	})
	t.Run("Empty_map", func(t *testing.T) {
		got := s71ScopeList(map[string]bool{})
		if got != "" {
			t.Errorf("s71ScopeList(empty) = %q, want empty", got)
		}
	})
}
