//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestXVM06_DesignSpecNil

package stml_design

import (
	"strings"
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
	"github.com/park-jun-woo/yongol/pkg/yongol"
)

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
