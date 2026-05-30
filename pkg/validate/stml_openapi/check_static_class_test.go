//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestCheckStaticClass — StaticElement 자체/Children class 검사

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestCheckStaticClass(t *testing.T) {
	// Minimal → no diagnostics.
	if got := checkStaticClass(stml.StaticElement{Tag: "div"}, "p.stml"); got != nil {
		t.Fatalf("empty static: expected nil, got %+v", got)
	}

	se := stml.StaticElement{
		Tag:       "section",
		ClassName: "wrap",
		Children: []stml.ChildNode{
			{Kind: "bind", Bind: &stml.FieldBind{Name: "c", ClassName: "p-1"}},
		},
	}
	diags := checkStaticClass(se, "p.stml")
	// static class(1) + child(1) = 2.
	if len(diags) != 2 {
		t.Fatalf("expected 2 diagnostics, got %d: %+v", len(diags), diags)
	}
}
