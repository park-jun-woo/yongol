//ff:func feature=validate type=test control=sequence topic=ssac-structural
//ff:what containsField — 필드 존재 여부 검증 (found/not found/empty)

package ssac

import "testing"

func TestContainsField(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		if !containsField([]string{"a", "b", "c"}, "b") {
			t.Error("expected true")
		}
	})
	t.Run("not found", func(t *testing.T) {
		if containsField([]string{"a", "b"}, "z") {
			t.Error("expected false")
		}
	})
	t.Run("empty slice", func(t *testing.T) {
		if containsField(nil, "a") {
			t.Error("expected false")
		}
	})
}
