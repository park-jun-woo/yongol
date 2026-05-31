//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXvm02RoundedUnit — xvm02Rounded rounded 토큰 정의 여부 분기 검증
package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXvm02RoundedUnit(t *testing.T) {
	fsWith := func() *yongol.Fullstack {
		return &yongol.Fullstack{DesignSpec: &design.DesignSpec{
			File:    "DESIGN.md",
			Rounded: map[string]string{"card": "0.5rem"},
		}}
	}

	t.Run("no design rounded returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{DesignSpec: &design.DesignSpec{Rounded: map[string]string{}}}
		if d := xvm02Rounded(fs, pageTokenRefs{}, overrideSet{}); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("defined token does not fire", func(t *testing.T) {
		tokens := pageTokenRefs{Rounded: []tokenRef{{File: "p.html", Class: "rounded-card", Name: "card"}}}
		if d := xvm02Rounded(fsWith(), tokens, overrideSet{}); len(d) != 0 {
			t.Errorf("expected no diagnostics, got %+v", d)
		}
	})

	t.Run("undefined token fires once despite duplicate", func(t *testing.T) {
		tokens := pageTokenRefs{Rounded: []tokenRef{
			{File: "p.html", Class: "rounded-ghost", Name: "ghost"},
			{File: "p.html", Class: "rounded-ghost", Name: "ghost"},
		}}
		d := xvm02Rounded(fsWith(), tokens, overrideSet{})
		if len(d) != 1 {
			t.Fatalf("expected 1 diagnostic (deduped), got %d: %+v", len(d), d)
		}
	})

	t.Run("overridden token skipped", func(t *testing.T) {
		tokens := pageTokenRefs{Rounded: []tokenRef{{File: "p.html", Class: "rounded-ghost", Name: "ghost"}}}
		ovr := overrideSet{"p.html": map[string]bool{"rounded-ghost": true}}
		if d := xvm02Rounded(fsWith(), tokens, ovr); len(d) != 0 {
			t.Errorf("expected no diagnostics when overridden, got %+v", d)
		}
	})
}
