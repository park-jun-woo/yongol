//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM04_Negative

package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM04_Negative(t *testing.T) {
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
				Static: &stml.StaticElement{ClassName: "font-Roboto"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm04Font(fs, tokens, nil)
	if len(diags) != 1 {
		t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XVM-04]") {
		t.Fatalf("expected [XVM-04], got %q", diags[0].Message)
	}
}
