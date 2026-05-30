//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM05_EmptyColors — DesignSpec.Colors 비어 있으면 early return nil

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM05_EmptyColors(t *testing.T) {
	fs := &yongol.Fullstack{
		SpecsDir: t.TempDir(),
		DesignSpec: &design.DesignSpec{
			File:   "DESIGN.md",
			Colors: map[string]string{},
		},
	}
	diags := xvm05Inline(fs, nil)
	if diags != nil {
		t.Fatalf("expected nil for empty colors, got %+v", diags)
	}
}
