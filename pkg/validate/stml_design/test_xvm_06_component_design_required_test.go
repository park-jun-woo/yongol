//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what XVM-06 테스트 — STML data-component의 DESIGN.md 정의 필수 검출

package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXVM06_Golden(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File: "DESIGN.md",
			Components: map[string]design.ComponentToken{
				"DatePicker": {Props: map[string]string{"variant": "outline"}},
			},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "ListItems",
				Components: []stml.ComponentRef{
					{Name: "DatePicker"},
				},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm06ComponentDesignRequired(fs, tokens)
	if len(diags) != 0 {
		t.Fatalf("expected 0 diags (all defined), got %d: %+v", len(diags), diags)
	}
}

func TestXVM06_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File: "DESIGN.md",
			Components: map[string]design.ComponentToken{
				"DatePicker": {Props: map[string]string{"variant": "outline"}},
			},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "ListItems",
				Components: []stml.ComponentRef{
					{Name: "DatePicker"},
					{Name: "Modal"},
				},
			}},
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm06ComponentDesignRequired(fs, tokens)
	if len(diags) != 1 {
		t.Fatalf("expected 1 (Modal missing), got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XVM-06]") {
		t.Fatalf("expected [XVM-06], got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "Modal") {
		t.Fatalf("expected mention of Modal, got %q", diags[0].Message)
	}
	if diags[0].File != "page.html" {
		t.Fatalf("expected file page.html, got %q", diags[0].File)
	}
}

func TestXVM06_DesignSpecNil(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: nil,
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
			Fetches: []stml.FetchBlock{{
				OperationID: "ListItems",
				Components: []stml.ComponentRef{
					{Name: "Modal"},
				},
			}},
		}},
	}
	// Run goes through the nil guard in Run(), but we test the function
	// directly with a nil DesignSpec — the defined map should be empty,
	// so every component is missing.
	tokens := pageTokenRefs{
		Components: []tokenRef{{File: "page.html", Name: "Modal"}},
	}
	// DesignSpec is nil — wrap in a Fullstack with nil DesignSpec
	fsNil := &yongol.Fullstack{DesignSpec: nil}
	diags := xvm06ComponentDesignRequired(fsNil, tokens)
	if len(diags) != 1 {
		t.Fatalf("expected 1 (Modal missing, DesignSpec nil), got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XVM-06]") {
		t.Fatalf("expected [XVM-06], got %q", diags[0].Message)
	}

	_ = fs // suppress unused warning
}

func TestXVM06_NoComponents(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File:       "DESIGN.md",
			Components: map[string]design.ComponentToken{},
		},
		STMLPages: []stml.PageSpec{{
			Name:     "page",
			FileName: "page.html",
		}},
	}
	tokens := extractAllTokens(fs)
	diags := xvm06ComponentDesignRequired(fs, tokens)
	if len(diags) != 0 {
		t.Fatalf("expected 0 (no STML components), got %d: %+v", len(diags), diags)
	}
}
