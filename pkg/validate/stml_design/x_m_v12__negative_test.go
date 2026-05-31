//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXMV12_Negative

package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMV12_Negative(t *testing.T) {
	fs := &yongol.Fullstack{
		DesignSpec: &design.DesignSpec{
			File: "DESIGN.md",
			Components: map[string]design.ComponentToken{
				"DatePicker": {Props: map[string]string{"variant": "outline"}},
				"Modal":      {Props: map[string]string{"size": "md"}},
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
	diags := xmv12DeadComponent(fs, tokens)
	if len(diags) != 1 {
		t.Fatalf("expected 1 (Modal unused), got %d: %+v", len(diags), diags)
	}
	if !strings.Contains(diags[0].Message, "[XMV-12]") {
		t.Fatalf("expected [XMV-12], got %q", diags[0].Message)
	}
	if !strings.Contains(diags[0].Message, "Modal") {
		t.Fatalf("expected mention of Modal, got %q", diags[0].Message)
	}
}
