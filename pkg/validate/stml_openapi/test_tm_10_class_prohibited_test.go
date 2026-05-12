//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM10_StaticElementWithClass_Error

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM10_StaticElementWithClass_Error(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Children: []stml.ChildNode{{
			Kind:   "static",
			Static: &stml.StaticElement{Tag: "div", ClassName: "bg-red-500"},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if !hasDiag(diags, "[TM-10]") {
		t.Errorf("expected TM-10 diagnostic for class on StaticElement, got %v", diags)
	}
}
