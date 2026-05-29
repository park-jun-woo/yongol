//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV07UnknownProp_NewFieldsInPropsNoWarning — 구조 필드가 props 아래 있어도 경고 없음

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
