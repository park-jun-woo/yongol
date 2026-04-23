//ff:func feature=validate type=test control=sequence topic=hurl-structural
//ff:what H-2 — scenario SSOT 가 없으면 규칙 침묵

package hurl

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

// TestH02EmptyTestsDirAbsent ensures the rule stays silent when scenario SSOT
// is absent (user opted out).
func TestH02EmptyTestsDirAbsent(t *testing.T) {
	fs := &yongol.Fullstack{}
	if diags := h02EmptyTestsDir(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostic when absent, got %d", len(diags))
	}
}
