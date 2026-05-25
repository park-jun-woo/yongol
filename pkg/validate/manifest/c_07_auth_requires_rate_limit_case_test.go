//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-structural
//ff:what runC07AuthRequiresRateLimit — TestC07AuthRequiresRateLimit table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runC07AuthRequiresRateLimit(t *testing.T, c TestC07AuthRequiresRateLimitCase) {
	t.Helper()
	diags := c07AuthRequiresRateLimit(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[C-7]") {
			t.Errorf("expected [C-7], got %q", d.Message)
		}
	}
}
