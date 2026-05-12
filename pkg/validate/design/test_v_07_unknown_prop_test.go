//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what V-07 테스트 — 미지의 component property

package design

import (
	"strings"
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/diagnostic"
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

func TestV07UnknownProp_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Components: map[string]pdesign.ComponentToken{
				"Card": {Props: map[string]string{
					"variant":      "solid",
					"customField":  "xyz",
				}},
			},
		},
	}
	got := v07UnknownProp(fs)
	if len(got) != 1 {
		t.Fatalf("expected 1 diagnostic, got %d: %+v", len(got), got)
	}
	if !strings.Contains(got[0].Message, "[V-07]") {
		t.Fatalf("message missing [V-07] prefix: %q", got[0].Message)
	}
	if got[0].Level != diagnostic.LevelWarning {
		t.Fatalf("V-07 should be WARNING, got %q", got[0].Level)
	}
}

func TestV07UnknownProp_NewFieldsInPropsNoWarning(t *testing.T) {
	// If a user accidentally puts base/variants/sizes/defaultVariant/defaultSize
	// under props: instead of as top-level component keys, V-07 should not warn.
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File: "DESIGN.md",
			Components: map[string]pdesign.ComponentToken{
				"Button": {Props: map[string]string{
					"base":           "inline-flex",
					"variants":       "primary | secondary",
					"sizes":          "sm | md",
					"defaultVariant": "primary",
					"defaultSize":    "md",
				}},
			},
		},
	}
	got := v07UnknownProp(fs)
	if len(got) != 0 {
		t.Fatalf("expected 0 diagnostics for new known props, got %d: %+v", len(got), got)
	}
}

func TestV07UnknownProp_NoComponents(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File:       "DESIGN.md",
			Components: nil,
		},
	}
	if got := v07UnknownProp(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
