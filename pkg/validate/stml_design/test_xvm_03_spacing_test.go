//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what XVM-03 테스트 — spacing 토큰 미정의 검출

package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM03_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File:    "DESIGN.md",
			Spacing: map[string]string{"lg": "1.5rem", "xl": "2rem"},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Children: []stml.ChildNode{{
				Kind:   "static",
				Static: &stml.StaticElement{ClassName: "p-lg gap-xl"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm03Spacing(fs, tokens, nil)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}
}

func TestXVM03_Negative(t *testing.T) {
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
				Static: &stml.StaticElement{ClassName: "m-huge"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm03Spacing(fs, tokens, nil)
	if len(diags) != 1 {
		t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XVM-03]") {
		t.Fatalf("expected [XVM-03], got %q", diags[0].Message)
	}
}

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
