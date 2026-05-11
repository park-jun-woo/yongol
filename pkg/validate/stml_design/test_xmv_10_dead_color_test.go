//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what XMV-10 테스트 — DESIGN.md colors 미참조 검출

package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMV10_Golden(t *testing.T) {
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
				Static: &stml.StaticElement{ClassName: "bg-primary"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xmv10DeadColor(fs, tokens)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}
}

func TestXMV10_Negative(t *testing.T) {
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
				Static: &stml.StaticElement{ClassName: "bg-primary"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xmv10DeadColor(fs, tokens)
	if len(diags) != 1 {
		t.Fatalf("expected 1 (accent unused), got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XMV-10]") {
		t.Fatalf("expected [XMV-10], got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "accent") {
		t.Fatalf("expected mention of accent, got %q", diags[0].Message)
	}
}
