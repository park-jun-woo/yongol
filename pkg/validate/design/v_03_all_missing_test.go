//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV03TypographyRequired_AllMissing — 모든 필수 필드 누락 시 V-03 진단 3건

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
