//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV07UnknownProp_Golden — 알려진 prop만 있을 때 진단 0건

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV07UnknownProp_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Components: map[string]pdesign.ComponentToken{
				"Button": {Props: map[string]string{
					"variant": "solid | outline",
					"size":    "sm | md | lg",
				}},
			},
		},
	}
	if got := v07UnknownProp(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
