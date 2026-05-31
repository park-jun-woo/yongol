//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM01_Golden

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM01_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File:   "DESIGN.md",
			Colors: map[string]string{"primary": "#6366F1", "accent": "#F59E0B"},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Children: []stml.ChildNode{{
				Kind:   "static",
				Static: &stml.StaticElement{ClassName: "bg-primary text-accent"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm01Color(fs, tokens, nil)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics, got %d: %+v", len(diags), diags)
	}
}
