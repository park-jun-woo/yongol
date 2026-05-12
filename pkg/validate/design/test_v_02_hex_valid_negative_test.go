//ff:func feature=validate type=test control=iteration dimension=1 topic=design-structural
//ff:what TestV02HexValid_Negative — 잘못된 hex color 시 V-02 진단 3건

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
