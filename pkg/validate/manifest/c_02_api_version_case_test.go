//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-structural
//ff:what runC02APIVersion — TestC02APIVersion table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runC02APIVersion(t *testing.T, c TestC02APIVersionCase) {
	t.Helper()
	diags := c02APIVersion(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d; diags=%v", len(diags), c.wantCount, diags)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[C-2]") {
			t.Errorf("expected [C-2], got %q", d.Message)
		}
	}
}
