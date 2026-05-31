//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM06_Negative

package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
