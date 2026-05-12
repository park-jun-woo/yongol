//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM11_LayoutFound_Pass

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM11_LayoutFound_Pass(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "dashboard-page",
		FileName: "dashboard-page.html",
		Layout:   "app",
	}}
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm11LayoutNotFound(pages, layouts)
	if hasDiag(diags, "[TM-11]") {
		t.Errorf("unexpected TM-11 for valid layout reference, got %v", diags)
	}
}
