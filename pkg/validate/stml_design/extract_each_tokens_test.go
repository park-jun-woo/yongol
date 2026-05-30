//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestExtractEachTokens — extractEachTokens EachBlock 토큰 참조 추출 검증

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestExtractEachTokens(t *testing.T) {
	eb := stml.EachBlock{
		ClassName:     "bg-primary",
		ItemClassName: "rounded-card",
		Binds:         []stml.FieldBind{{ClassName: "p-section"}},
		Components:    []stml.ComponentRef{{Name: "DatePicker", ClassName: "font-brand"}},
		Children: []stml.ChildNode{
			{Kind: "static", Static: &stml.StaticElement{ClassName: "m-gutter"}},
		},
	}

	var out pageTokenRefs
	extractEachTokens(eb, "p.stml", &out)

	if len(out.Colors) != 1 {
		t.Errorf("expected 1 color, got %+v", out.Colors)
	}
	if len(out.Rounded) != 1 {
		t.Errorf("expected 1 rounded, got %+v", out.Rounded)
	}
	if len(out.Spacing) != 2 {
		t.Errorf("expected 2 spacing (bind + child), got %+v", out.Spacing)
	}
	if len(out.Fonts) != 1 {
		t.Errorf("expected 1 font, got %+v", out.Fonts)
	}
	if len(out.Components) != 1 {
		t.Errorf("expected 1 component, got %+v", out.Components)
	}
}
