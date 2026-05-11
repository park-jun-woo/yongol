//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what V-05 테스트 — 토큰 참조 resolve

package design

import (
	"strings"
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

func TestV05TokenRefResolve_NoRefNoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Components: map[string]pdesign.ComponentToken{
				"Button": {Props: map[string]string{
					"variant": "solid | outline | ghost",
				}},
			},
		},
	}
	if got := v05TokenRefResolve(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
