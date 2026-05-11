//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what XVM-02 테스트 — rounded 토큰 미정의 검출

package stml_design

import (
	"strings"
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

func TestXVM02_Negative(t *testing.T) {
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
				Static: &stml.StaticElement{ClassName: "rounded-mega"},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm02Rounded(fs, tokens, nil)
	if len(diags) != 1 {
		t.Fatalf("expected 1, got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XVM-02]") {
		t.Fatalf("expected [XVM-02], got %q", diags[0].Message)
	}
}
