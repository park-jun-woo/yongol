//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXMV11_Golden

package stml_design

import (
	"testing"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMV11_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File: "DESIGN.md",
			Typography: map[string]design.TypographyToken{
				"heading": {FontFamily: "Inter"},
			},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Children: []stml.ChildNode{{
				Kind:   "static",
				Static: &stml.StaticElement{ClassName: "font-Inter"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xmv11DeadTypography(fs, tokens)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}
}
