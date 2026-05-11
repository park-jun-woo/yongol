//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what XVM-01 테스트 — 색상 토큰 미정의 검출

package stml_design

import (
	"strings"
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

func TestXVM01_SkipsNumericAndBuiltin(t *testing.T) {
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
				Static: &stml.StaticElement{ClassName: "bg-white text-black border-transparent bg-[#123456]"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm01Color(fs, tokens, nil)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diagnostics (builtins skipped), got %d: %+v", len(diags), diags)
	}
}

func TestXVM01_SkipsTailwindPalette(t *testing.T) {
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
				Static: &stml.StaticElement{ClassName: "text-gray-500 bg-red-200"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm01Color(fs, tokens, nil)
	if len(diags) != 0 {
		t.Fatalf("expected 0 (palette skipped), got %d: %+v", len(diags), diags)
	}
}
