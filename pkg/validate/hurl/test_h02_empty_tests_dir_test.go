//ff:func feature=validate type=test control=sequence topic=hurl-structural
//ff:what H-2 positive — tests/ 디렉토리 SSOTDeclared 상태일 때 WARNING

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
