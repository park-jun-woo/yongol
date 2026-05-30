//ff:func feature=gen-fastapi type=test control=sequence
//ff:what TestRenderInlineStubs — same-feature @call inline stub 함수 렌더링

package ssac

import (
	"strings"
	"testing"
)

func TestRenderInlineStubs(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		if got := renderInlineStubs("billing", nil); got != "" {
			t.Errorf("expected empty, got %q", got)
		}
	})
	t.Run("Stubs", func(t *testing.T) {
		got := renderInlineStubs("billing", []string{"calc_total", "apply_discount"})
		if !strings.Contains(got, "async def calc_total(*args, **kwargs):") {
			t.Errorf("missing calc_total stub: %q", got)
		}
		if !strings.Contains(got, `raise NotImplementedError("billing.apply_discount not implemented")`) {
			t.Errorf("missing apply_discount NotImplementedError: %q", got)
		}
		if !strings.HasPrefix(got, "\n# --- same-feature @call stubs") {
			t.Errorf("missing header: %q", got)
		}
	})
}
