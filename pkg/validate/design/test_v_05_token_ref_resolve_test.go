//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV05TokenRefResolve_Golden — 정상 참조 시 진단 0건

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV05TokenRefResolve_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Colors: map[string]string{
				"primary": "#6366F1",
			},
			Spacing: map[string]string{
				"md": "1rem",
			},
			Components: map[string]pdesign.ComponentToken{
				"Button": {Props: map[string]string{
					"bg":      "{colors.primary}",
					"padding": "{spacing.md}",
				}},
			},
		},
	}
	if got := v05TokenRefResolve(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
