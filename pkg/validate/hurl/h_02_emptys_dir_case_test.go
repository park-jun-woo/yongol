//ff:func feature=validate type=test-helper control=iteration dimension=1 topic=hurl-structural
//ff:what runH02EmptyTestsDir — TestH02EmptyTestsDir table-driven 개별 케이스 검증

package hurl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/diagnostic"
)

func runH02EmptyTestsDir(t *testing.T, c TestH02EmptyTestsDirCase) {
	t.Helper()
	diags := h02EmptyTestsDir(c.fs)
	if len(diags) != c.wantCount {
		t.Fatalf("got %d diags, want %d", len(diags), c.wantCount)
	}
	for _, d := range diags {
		if d.Level != diagnostic.LevelWarning {
			t.Errorf("expected LevelWarning, got %q", d.Level)
		}
		if !strings.Contains(d.Message, "[H-2]") {
			t.Errorf("message should contain [H-2], got %q", d.Message)
		}
	}
}
