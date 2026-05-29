//ff:func feature=validate type=test control=iteration dimension=1 topic=design-structural
//ff:what TestV04DimensionValid_Negative — 잘못된 dimension 시 V-04 진단 2건

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
