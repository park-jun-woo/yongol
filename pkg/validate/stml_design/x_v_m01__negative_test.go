//ff:func feature=validate type=test control=iteration dimension=1 topic=stml-design
//ff:what TestXVM01_Negative

package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM01_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File:   "DESIGN.md",
			Colors: map[string]string{"primary": "#6366F1"},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Children: []stml.ChildNode{{
				Kind:   "static",
				Static: &stml.StaticElement{ClassName: "bg-missing text-unknown"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm01Color(fs, tokens, nil)
	if len(diags) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d: %+v", len(diags), diags)
	}
	for _, d := range diags {
		if !strings.Contains(d.Message, "[XVM-01]") {
			t.Fatalf("expected [XVM-01] prefix, got %q", d.Message)
		}
	}
}
