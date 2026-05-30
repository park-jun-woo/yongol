//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestExtractPageTokens — extractPageTokens 단일 STML 페이지 토큰 참조 추출 검증

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestExtractPageTokens(t *testing.T) {
	page := stml.PageSpec{
		FileName: "p.stml",
		Fetches:  []stml.FetchBlock{{ClassName: "bg-primary"}},
		Actions:  []stml.ActionBlock{{ClassName: "rounded-card"}},
		Children: []stml.ChildNode{
			{Kind: "static", Static: &stml.StaticElement{ClassName: "p-section"}},
		},
	}

	var out pageTokenRefs
	extractPageTokens(page, &out)

	if len(out.Colors) != 1 {
		t.Errorf("expected 1 color from fetch, got %+v", out.Colors)
	}
	if len(out.Rounded) != 1 {
		t.Errorf("expected 1 rounded from action, got %+v", out.Rounded)
	}
	if len(out.Spacing) != 1 {
		t.Errorf("expected 1 spacing from child, got %+v", out.Spacing)
	}
}
