//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXMV11_Negative

package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMV11_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File: "DESIGN.md",
			Typography: map[string]design.TypographyToken{
				"heading": {FontFamily: "Inter"},
				"body":    {FontFamily: "Roboto"},
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
	if len(diags) != 1 {
		t.Fatalf("expected 1 (Roboto unused), got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XMV-11]") {
		t.Fatalf("expected [XMV-11], got %q", diags[0].Message)
	}
}
