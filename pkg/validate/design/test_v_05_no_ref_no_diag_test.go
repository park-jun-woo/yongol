//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV05TokenRefResolve_NoRefNoDiag — 참조 없는 prop 시 진단 0건

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV05TokenRefResolve_NoRefNoDiag(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Components: map[string]pdesign.ComponentToken{
				"Button": {Props: map[string]string{
					"variant": "solid | outline | ghost",
				}},
			},
		},
	}
	if got := v05TokenRefResolve(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
