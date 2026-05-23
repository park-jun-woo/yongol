//ff:func feature=validate type=test control=sequence topic=ssac-openapi
//ff:what toSet — nil/empty/복수 항목 변환 검증

package openapi_ssac

import "testing"

func TestToSet(t *testing.T) {
	t.Run("nil returns empty", func(t *testing.T) {
		got := toSet(nil)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("converts slice to set", func(t *testing.T) {
		got := toSet([]string{"a", "b", "c"})
		if len(got) != 3 || !got["a"] || !got["b"] || !got["c"] {
			t.Errorf("unexpected set: %v", got)
		}
	})
}
