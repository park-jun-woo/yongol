//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV02HexValid_Golden — 유효한 hex color 시 진단 0건

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV02HexValid_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Colors: map[string]string{
				"primary": "#6366F1",
				"short":   "#FFF",
				"with4":   "#FFFA",
				"with8":   "#6366F1AA",
			},
		},
	}
	if got := v02HexValid(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
