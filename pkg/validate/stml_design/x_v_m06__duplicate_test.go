//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM06_Duplicate — 동일 미정의 컴포넌트 중복 참조는 1회만 진단

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM06_Duplicate(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File:       "DESIGN.md",
			Components: map[string]design.ComponentToken{},
		},
	}
	tokens := pageTokenRefs{
		Components: []tokenRef{
			{File: "page.html", Name: "Modal"},
			{File: "other.html", Name: "Modal"},
			{File: "page.html", Name: "Accordion"},
		},
	}
	diags := xvm06ComponentDesignRequired(fs, tokens)
	if len(diags) != 2 {
		t.Fatalf("expected 2 (dedup Modal + Accordion), got %d: %+v", len(diags), diags)
	}
	// Sorted by name: Accordion before Modal.
	if diags[0].Message == "" || diags[1].Message == "" {
		t.Fatalf("missing message: %+v", diags)
	}
}
