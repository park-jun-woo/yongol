//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM10_MultipleViolations_CountsAll

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM10_MultipleViolations_CountsAll(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Children: []stml.ChildNode{
			{Kind: "static", Static: &stml.StaticElement{Tag: "div", ClassName: "bg-red"}},
			{Kind: "static", Static: &stml.StaticElement{Tag: "span", ClassName: "text-lg"}},
			{Kind: "static", Static: &stml.StaticElement{Tag: "p"}},
		},
	}}
	diags := tm10ClassProhibited(pages)
	count := countDiag(diags, "[TM-10]")
	if count != 2 {
		t.Errorf("expected 2 TM-10 diagnostics, got %d: %+v", count, diags)
	}
}
