//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV07UnknownProp_NoComponents — components nil 시 진단 0건

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
