//ff:func feature=validate type=test control=selection topic=stml-design
//ff:what TestXvm04FontUnit — xvm04Font font 토큰의 typography fontFamily 매칭 분기 검증

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXvm04FontUnit(t *testing.T) {
	fsWith := func() *yongol.Fullstack {
		return &yongol.Fullstack{DesignSpec: &design.DesignSpec{
			File: "DESIGN.md",
			Typography: map[string]design.TypographyToken{
				"heading": {FontFamily: "Inter"},
			},
		}}
	}

	t.Run("no typography returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{DesignSpec: &design.DesignSpec{Typography: map[string]design.TypographyToken{}}}
		if d := xvm04Font(fs, pageTokenRefs{}, overrideSet{}); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("typography present but no fontFamily returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{DesignSpec: &design.DesignSpec{
			Typography: map[string]design.TypographyToken{"blank": {FontFamily: ""}},
		}}
		if d := xvm04Font(fs, pageTokenRefs{}, overrideSet{}); d != nil {
			t.Errorf("expected nil when no known fonts, got %+v", d)
		}
	})

	t.Run("known font (case-insensitive) does not fire", func(t *testing.T) {
		tokens := pageTokenRefs{Fonts: []tokenRef{{File: "p.html", Class: "font-inter", Name: "inter"}}}
		if d := xvm04Font(fsWith(), tokens, overrideSet{}); len(d) != 0 {
			t.Errorf("expected no diagnostics, got %+v", d)
		}
	})

	t.Run("unknown font fires once despite duplicate", func(t *testing.T) {
		tokens := pageTokenRefs{Fonts: []tokenRef{
			{File: "p.html", Class: "font-ghost", Name: "ghost"},
			{File: "p.html", Class: "font-ghost", Name: "ghost"},
		}}
		d := xvm04Font(fsWith(), tokens, overrideSet{})
		if len(d) != 1 {
			t.Fatalf("expected 1 diagnostic (deduped), got %d: %+v", len(d), d)
		}
	})

	t.Run("overridden font skipped", func(t *testing.T) {
		tokens := pageTokenRefs{Fonts: []tokenRef{{File: "p.html", Class: "font-ghost", Name: "ghost"}}}
		ovr := overrideSet{"p.html": map[string]bool{"font-ghost": true}}
		if d := xvm04Font(fsWith(), tokens, ovr); len(d) != 0 {
			t.Errorf("expected no diagnostics when overridden, got %+v", d)
		}
	})
}
