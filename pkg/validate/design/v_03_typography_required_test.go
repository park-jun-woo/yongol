//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV03TypographyRequired_Golden — 모든 필드 있을 때 진단 0건

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV03TypographyRequired_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Typography: map[string]pdesign.TypographyToken{
				"heading": {FontFamily: "Inter", FontSize: "1.5rem", FontWeight: "700"},
				"body":    {FontFamily: "Inter", FontSize: "1rem", FontWeight: "400"},
			},
		},
	}
	if got := v03TypographyRequired(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
