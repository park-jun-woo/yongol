//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM13_AllUsed_Pass

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM13_AllUsed_Pass(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "dashboard-page", FileName: "dashboard-page.html", Layout: "app"},
		{Name: "login-page", FileName: "login-page.html", Layout: "auth"},
	}
	layouts := []stml.LayoutSpec{
		{Name: "app", File: "layouts/app.html"},
		{Name: "auth", File: "layouts/auth.html"},
	}
	diags := tm13UnusedLayout(pages, layouts, "", nil)
	if hasDiag(diags, "[TM-13]") {
		t.Errorf("unexpected TM-13 when all layouts are used, got %v", diags)
	}
}
