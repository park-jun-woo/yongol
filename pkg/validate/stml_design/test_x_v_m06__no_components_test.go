//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM06_NoComponents

package stml_design

import (
	"testing"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM06_NoComponents(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File:       "DESIGN.md",
			Components: map[string]design.ComponentToken{},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm06ComponentDesignRequired(fs, tokens)
	if len(diags) != 0 {
		t.Fatalf("expected 0 (no STML components), got %d: %+v", len(diags), diags)
	}
}
