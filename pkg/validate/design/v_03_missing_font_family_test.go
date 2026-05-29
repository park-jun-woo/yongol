//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV03TypographyRequired_MissingFontFamily — fontFamily 누락 시 V-03 진단 1건

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV03TypographyRequired_MissingFontFamily(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Typography: map[string]pdesign.TypographyToken{
				"heading": {FontFamily: "", FontSize: "1.5rem", FontWeight: "700"},
			},
		},
	}
	got := v03TypographyRequired(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[V-03]") || !strings.Contains(got[0].Message, "fontFamily") {
		t.Fatalf("unexpected message: %q", got[0].Message)
	}
}
