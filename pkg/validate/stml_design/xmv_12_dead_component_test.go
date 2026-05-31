//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXMV12_Golden

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/design"
	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

func TestXMV12_Golden(t *testing.T) {
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
	diags := xmv12DeadComponent(fs, tokens)
	if len(diags) != 0 {
		t.Fatalf("expected 0, got %d: %+v", len(diags), diags)
	}
}
