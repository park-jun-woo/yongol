//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what V-02 테스트 — hex color 유효성

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV02HexValid_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Colors: map[string]string{
				"primary":   "#6366F1",
				"short":     "#FFF",
				"with4":     "#FFFA",
				"with8":     "#6366F1AA",
			},
		},
	}
	if got := v02HexValid(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}

func TestV02HexValid_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Colors: map[string]string{
				"bad1": "red",
				"bad2": "#GGG",
				"bad3": "#12345",
			},
		},
	}
	got := v02HexValid(fs)
	if len(got) != 3 {
		t.Fatalf("expected 3 diagnostics, got %d: %+v", len(got), got)
	}
	for _, d := range got {
		if !strings.Contains(d.Message, "[V-02]") {
			t.Fatalf("message missing [V-02] prefix: %q", d.Message)
		}
	}
}
