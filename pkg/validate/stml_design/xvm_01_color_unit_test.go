//ff:func feature=validate type=test control=selection topic=stml-design
//ff:what TestXvm01ColorUnit — xvm01Color 색상 토큰 정의 여부 분기 검증

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXvm01ColorUnit(t *testing.T) {
	fsWith := func() *yongol.Fullstack {
		return &yongol.Fullstack{DesignSpec: &design.DesignSpec{
			File:   "DESIGN.md",
			Colors: map[string]string{"primary": "#6366F1"},
		}}
	}

	t.Run("no design colors returns nil", func(t *testing.T) {
		fs := &yongol.Fullstack{DesignSpec: &design.DesignSpec{Colors: map[string]string{}}}
		if d := xvm01Color(fs, pageTokenRefs{}, overrideSet{}); d != nil {
			t.Errorf("expected nil, got %+v", d)
		}
	})

	t.Run("defined token does not fire", func(t *testing.T) {
		tokens := pageTokenRefs{Colors: []tokenRef{{File: "p.html", Class: "bg-primary", Name: "primary"}}}
		if d := xvm01Color(fsWith(), tokens, overrideSet{}); len(d) != 0 {
			t.Errorf("expected no diagnostics, got %+v", d)
		}
	})

	t.Run("undefined token fires once despite duplicate", func(t *testing.T) {
		tokens := pageTokenRefs{Colors: []tokenRef{
			{File: "p.html", Class: "bg-ghost", Name: "ghost"},
			{File: "p.html", Class: "text-ghost", Name: "ghost"},
		}}
		d := xvm01Color(fsWith(), tokens, overrideSet{})
		if len(d) != 1 {
			t.Fatalf("expected 1 diagnostic (deduped), got %d: %+v", len(d), d)
		}
	})

	t.Run("overridden token skipped", func(t *testing.T) {
		tokens := pageTokenRefs{Colors: []tokenRef{{File: "p.html", Class: "bg-ghost", Name: "ghost"}}}
		ovr := overrideSet{"p.html": map[string]bool{"bg-ghost": true}}
		if d := xvm01Color(fsWith(), tokens, ovr); len(d) != 0 {
			t.Errorf("expected no diagnostics when overridden, got %+v", d)
		}
	})
}
