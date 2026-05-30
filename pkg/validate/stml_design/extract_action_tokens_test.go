//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestExtractActionTokens — extractActionTokens ActionBlock 토큰 참조 추출 검증

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestExtractActionTokens(t *testing.T) {
	ab := stml.ActionBlock{
		ClassName: "bg-primary",
		Fields: []stml.FieldBind{
			{Name: "Status", ClassName: "rounded-card"},
		},
		Children: []stml.ChildNode{
			{Kind: "static", Static: &stml.StaticElement{Tag: "div", ClassName: "p-section"}},
		},
	}

	var out pageTokenRefs
	extractActionTokens(ab, "p.stml", &out)

	if len(out.Colors) != 1 {
		t.Errorf("expected 1 color from ClassName, got %+v", out.Colors)
	}
	if len(out.Rounded) != 1 {
		t.Errorf("expected 1 rounded from field, got %+v", out.Rounded)
	}
	if len(out.Spacing) != 1 {
		t.Errorf("expected 1 spacing from child, got %+v", out.Spacing)
	}
}
