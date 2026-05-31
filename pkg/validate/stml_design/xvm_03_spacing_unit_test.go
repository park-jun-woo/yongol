//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXvm03SpacingUnit — xvm03Spacing spacing 토큰 정의 여부 분기 검증
package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXvm03SpacingUnit(t *testing.T) {
	fsWith := func() *yongol.Fullstack {
		return &yongol.Fullstack{DesignSpec: &design.DesignSpec{
			File:    "DESIGN.md",
			Spacing: map[string]string{"section": "2rem"},
		}}
	}

	t.Run("no design spacing returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{DesignSpec: &design.DesignSpec{Spacing: map[string]string{}}}
		if d := xvm03Spacing(fs, pageTokenRefs{}, overrideSet{}); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("defined token does not fire", func(t *testing.T) {
		tokens := pageTokenRefs{Spacing: []tokenRef{{File: "p.html", Class: "p-section", Name: "section"}}}
		if d := xvm03Spacing(fsWith(), tokens, overrideSet{}); len(d) != 0 {
			t.Errorf("expected no diagnostics, got %+v", d)
		}
	})

	t.Run("undefined token fires once despite duplicate", func(t *testing.T) {
		tokens := pageTokenRefs{Spacing: []tokenRef{
			{File: "p.html", Class: "p-ghost", Name: "ghost"},
			{File: "p.html", Class: "m-ghost", Name: "ghost"},
		}}
		d := xvm03Spacing(fsWith(), tokens, overrideSet{})
		if len(d) != 1 {
			t.Fatalf("expected 1 diagnostic (deduped), got %d: %+v", len(d), d)
		}
	})

	t.Run("overridden token skipped", func(t *testing.T) {
		tokens := pageTokenRefs{Spacing: []tokenRef{{File: "p.html", Class: "p-ghost", Name: "ghost"}}}
		ovr := overrideSet{"p.html": map[string]bool{"p-ghost": true}}
		if d := xvm03Spacing(fsWith(), tokens, ovr); len(d) != 0 {
			t.Errorf("expected no diagnostics when overridden, got %+v", d)
		}
	})
}
