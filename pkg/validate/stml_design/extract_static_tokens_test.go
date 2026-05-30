//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestExtractStaticTokens — extractStaticTokens StaticElement 재귀 토큰 추출 검증

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestExtractStaticTokens(t *testing.T) {
	se := stml.StaticElement{
		ClassName: "bg-primary",
		Children: []stml.ChildNode{
			{Kind: "static", Static: &stml.StaticElement{ClassName: "p-section"}},
		},
	}

	var out pageTokenRefs
	extractStaticTokens(se, "p.stml", &out)

	if len(out.Colors) != 1 {
		t.Errorf("expected 1 color, got %+v", out.Colors)
	}
	if len(out.Spacing) != 1 {
		t.Errorf("expected 1 spacing from nested child, got %+v", out.Spacing)
	}
}
