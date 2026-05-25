//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-structural
//ff:what runC08RateLimitLoginRequired — TestC08RateLimitLoginRequired table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runC08RateLimitLoginRequired(t *testing.T, c TestC08RateLimitLoginRequiredCase) {
	t.Helper()
	diags := c08RateLimitLoginRequired(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[C-8]") {
			t.Errorf("expected [C-8], got %q", d.Message)
		}
	}
}
