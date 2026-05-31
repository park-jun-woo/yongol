//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM04_CaseInsensitive

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM04_CaseInsensitive(t *testing.T) {
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
				Static: &stml.StaticElement{ClassName: "font-inter"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm04Font(fs, tokens, nil)
	if len(diags) != 0 {
		t.Fatalf("expected 0 (case-insensitive match), got %d: %+v", len(diags), diags)
	}
}
