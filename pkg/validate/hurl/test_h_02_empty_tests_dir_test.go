//ff:func feature=validate type=test control=sequence topic=hurl-structural
//ff:what H-2 테스트 — tests/ 디렉토리 SSOTDeclared 상태일 때 WARNING 발화

package hurl

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestH02EmptyTestsDir(t *testing.T) {
	fs := &yongol.Fullstack{
		Presences: map[yongol.SSOTKind]yongol.SSOTPresence{
			yongol.KindScenario: yongol.SSOTDeclared,
		},
	}
	diags := h02EmptyTestsDir(fs)
	if len(diags) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d", len(diags))
	}
	if !strings.Contains(diags[0].Message, "[H-2]") {
		t.Errorf("rule id missing: %q", diags[0].Message)
	}
	if diags[0].Level != "WARNING" {
		t.Errorf("want WARNING, got %s", diags[0].Level)
	}
}

// TestH02EmptyTestsDirAbsent ensures the rule stays silent when scenario SSOT
// is absent (user opted out).
func TestH02EmptyTestsDirAbsent(t *testing.T) {
	fs := &yongol.Fullstack{}
	if diags := h02EmptyTestsDir(fs); len(diags) != 0 {
		t.Fatalf("expected 0 diagnostic when absent, got %d", len(diags))
	}
}
