//ff:func feature=validate type=test control=sequence topic=config-check
//ff:what middlewareSet — nil/empty/복수 항목 set 변환 검증

package openapi_manifest

import "testing"

func TestMiddlewareSet(t *testing.T) {
	t.Run("nil returns empty", func(t *testing.T) {
		got := middlewareSet(nil)
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("empty returns empty", func(t *testing.T) {
		got := middlewareSet([]string{})
		if len(got) != 0 {
			t.Errorf("expected empty, got %v", got)
		}
	})

	t.Run("multiple items", func(t *testing.T) {
		got := middlewareSet([]string{"cors", "auth", "logger"})
		if len(got) != 3 {
			t.Fatalf("expected 3, got %d", len(got))
		}
		if !got["cors"] || !got["auth"] || !got["logger"] {
			t.Errorf("missing expected items: %v", got)
		}
	})
}
