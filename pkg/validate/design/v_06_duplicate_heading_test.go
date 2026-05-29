//ff:func feature=validate type=test control=sequence topic=design-structural
//ff:what TestV06DuplicateHeading_Golden — 중복 없을 때 진단 0건

package design

import (
	"testing"

	pdesign "github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestV06DuplicateHeading_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &pdesign.DesignSpec{
			File:     "DESIGN.md",
			Headings: []string{"Colors", "Typography", "Components"},
		},
	}
	if got := v06DuplicateHeading(fs); len(got) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(got), got)
	}
}
