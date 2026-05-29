//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV05TokenRefResolve_Negative — 미해석 참조 시 V-05 진단 1건

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV05TokenRefResolve_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File:   "DESIGN.md",
			Colors: map[string]string{},
			Components: map[string]pdesign.ComponentToken{
				"Card": {Props: map[string]string{
					"bg": "{colors.missing}",
				}},
			},
		},
	}
	got := v05TokenRefResolve(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[V-05]") {
		t.Fatalf("message missing [V-05] prefix: %q", got[0].Message)
	}
	if !strings.Contains(got[0].Message, "colors.missing") {
		t.Fatalf("message should reference the unresolved token: %q", got[0].Message)
	}
}
