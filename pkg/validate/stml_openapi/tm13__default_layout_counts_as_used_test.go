//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM13_DefaultLayoutCountsAsUsed

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM13_DefaultLayoutCountsAsUsed(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "login-page",
		FileName: "login-page.html",
		Layout:   "auth",
	}}
	layouts := []stml.LayoutSpec{
		{Name: "app", File: "layouts/app.html"},
		{Name: "auth", File: "layouts/auth.html"},
	}
	// "app" is only used as defaultLayout, not by any page directly
	diags := tm13UnusedLayout(pages, layouts, "app", nil)
	if hasDiag(diags, "[TM-13]") {
		t.Errorf("unexpected TM-13 when layout is defaultLayout, got %v", diags)
	}
}
