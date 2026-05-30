//ff:func feature=validate type=test control=sequence topic=stml-design
//ff:what TestExtractFetchTokens — extractFetchTokens FetchBlock 토큰 참조 추출 검증

package stml_design

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestExtractFetchTokens(t *testing.T) {
	fb := stml.FetchBlock{
		ClassName: "bg-primary",
		Binds:     []stml.FieldBind{{ClassName: "rounded-card"}},
		Eaches:    []stml.EachBlock{{ClassName: "p-section"}},
		Components: []stml.ComponentRef{
			{Name: "DatePicker", ClassName: "font-brand"},
		},
		Children: []stml.ChildNode{
			{Kind: "static", Static: &stml.StaticElement{ClassName: "m-gutter"}},
		},
		NestedFetches: []stml.FetchBlock{
			{ClassName: "text-accent"},
		},
	}

	var out pageTokenRefs
	extractFetchTokens(fb, "p.stml", &out)

	if len(out.Colors) != 2 {
		t.Errorf("expected 2 colors (root + nested), got %+v", out.Colors)
	}
	if len(out.Rounded) != 1 {
		t.Errorf("expected 1 rounded, got %+v", out.Rounded)
	}
	if len(out.Spacing) != 2 {
		t.Errorf("expected 2 spacing (each + child), got %+v", out.Spacing)
	}
	if len(out.Fonts) != 1 {
		t.Errorf("expected 1 font, got %+v", out.Fonts)
	}
	if len(out.Components) != 1 {
		t.Errorf("expected 1 component, got %+v", out.Components)
	}
}
