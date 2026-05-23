//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=manifest-structural
//ff:what runC06BackendAuthRequired — TestC06BackendAuthRequired table-driven 개별 케이스 검증

package manifest

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runC06BackendAuthRequired(t *testing.T, c TestC06BackendAuthRequiredCase) {
	t.Helper()
	diags := c06BackendAuthRequired(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelError {
			t.Errorf("expected LevelError, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[C-6]") {
			t.Errorf("expected [C-6], got %q", d.Message)
		}
	}
}
