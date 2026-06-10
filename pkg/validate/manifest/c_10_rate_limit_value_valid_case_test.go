//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-structural
//ff:what runC10RateLimitValueValid — TestC10RateLimitValueValid table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runC10RateLimitValueValid(t *testing.T, c TestC10RateLimitValueValidCase) {
	t.Helper()
	diags := c10RateLimitValueValid(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if d.Phase != diagnostic.PhaseValidate {
			t.Errorf("expected PhaseValidate, got %q", d.Phase)
		}
		if !strings.Contains(d.Message, "[C-10]") {
			t.Errorf("expected [C-10], got %q", d.Message)
		}
	}
}
