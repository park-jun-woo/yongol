//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV04DimensionValid_Golden — 유효한 dimension 시 진단 0건

package design

import (
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
