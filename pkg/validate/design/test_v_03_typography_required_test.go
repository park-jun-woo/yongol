//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what V-03 테스트 — typography 필수 필드

package design

import (
	"strings"
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

func TestV03TypographyRequired_AllMissing(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Typography: map[string]pdesign.TypographyToken{
				"broken": {},
			},
		},
	}
	got := v03TypographyRequired(fs)
	if len(got) != 3 {
		t.Fatalf("expected 3 diagnostics, got %d: %+v", len(got), got)
	}
}
