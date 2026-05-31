//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXmv11DeadTypographyBranches — xmv11DeadTypography early-return / empty fontFamily skip 검증
package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXmv11DeadTypographyBranches(t *testing.T) {
	t.Run("no typography returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DesignSpec: &design.DesignSpec{File: "DESIGN.md", Typography: map[string]design.TypographyToken{}},
		}
		if d := xmv11DeadTypography(fs, pageTokenRefs{}); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("empty fontFamily skipped", func(t *testing.T) {
		fs := &yongol.Fullstack{
			DesignSpec: &design.DesignSpec{
				File: "DESIGN.md",
				Typography: map[string]design.TypographyToken{
					"blank": {FontFamily: ""},
				},
			},
		}
		if d := xmv11DeadTypography(fs, pageTokenRefs{}); len(d) != 0 {
			t.Errorf("expected no diagnostics for empty fontFamily, got %+v", d)
		}
	})
}
