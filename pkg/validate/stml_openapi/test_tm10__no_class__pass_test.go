//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM10_NoClass_Pass

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM10_NoClass_Pass(t *testing.T) {
	pages := []stml.PageSpec{{
		FileName: "page.html",
		Children: []stml.ChildNode{{
			Kind:   "static",
			Static: &stml.StaticElement{Tag: "div"},
		}},
	}}
	diags := tm10ClassProhibited(pages)
	if hasDiag(diags, "[TM-10]") {
		t.Errorf("unexpected TM-10 diagnostic for element without class, got %v", diags)
	}
}
