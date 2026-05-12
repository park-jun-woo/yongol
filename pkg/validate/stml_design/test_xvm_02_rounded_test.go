//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM02_Golden

package stml_design

import (
	"testing"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM02_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File:    "DESIGN.md",
			Rounded: map[string]string{"card": "0.5rem"},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Children: []stml.ChildNode{{
				Kind:   "static",
				Static: &stml.StaticElement{ClassName: "rounded-card"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm02Rounded(fs, tokens, nil)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}
}
