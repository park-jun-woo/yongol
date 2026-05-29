//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM11_EmptyLayout_Skip

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM11_EmptyLayout_Skip(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "login-page",
		FileName: "login-page.html",
		Layout:   "",
	}}
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm11LayoutNotFound(pages, layouts)
	if hasDiag(diags, "[TM-11]") {
		t.Errorf("unexpected TM-11 for page with empty layout, got %v", diags)
	}
}
