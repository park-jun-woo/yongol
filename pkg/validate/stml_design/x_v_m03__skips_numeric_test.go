//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM03_SkipsNumeric

package stml_design

import (
	"testing"
	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM03_SkipsNumeric(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File:    "DESIGN.md",
			Spacing: map[string]string{"lg": "1.5rem"},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Children: []stml.ChildNode{{
				Kind:   "static",
				Static: &stml.StaticElement{ClassName: "p-4 m-0.5 gap-1/2"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm03Spacing(fs, tokens, nil)
	if len(diags) != 0 {
		t.Fatalf("expected 0 (numeric skipped), got %d: %+v", len(diags), diags)
	}
}
