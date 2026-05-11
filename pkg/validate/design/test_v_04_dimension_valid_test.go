//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what V-04 테스트 — dimension 유효성

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV04DimensionValid_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Rounded: map[string]string{
				"sm": "0.25rem",
				"md": "4px",
				"lg": "8",
			},
			Spacing: map[string]string{
				"xs": "0.25em",
				"sm": "0.5rem",
				"md": "16px",
				"lg": "32",
			},
		},
	}
	if got := v04DimensionValid(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}

func TestV04DimensionValid_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Rounded: map[string]string{
				"bad": "large",
			},
			Spacing: map[string]string{
				"bad": "auto",
			},
		},
	}
	got := v04DimensionValid(fs)
	if len(got) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d: %+v", len(got), got)
	}
	for _, d := range got {
		if !strings.Contains(d.Message, "[V-04]") {
			t.Fatalf("message missing [V-04] prefix: %q", d.Message)
		}
	}
}
