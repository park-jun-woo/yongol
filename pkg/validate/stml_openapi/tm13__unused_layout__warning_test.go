//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM13_UnusedLayout_Warning

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM13_UnusedLayout_Warning(t *testing.T) {
	pages := []stml.PageSpec{{
		Name:     "dashboard-page",
		FileName: "dashboard-page.html",
		Layout:   "app",
	}}
	layouts := []stml.LayoutSpec{
		{Name: "app", File: "layouts/app.html"},
		{Name: "auth", File: "layouts/auth.html"},
	}
	diags := tm13UnusedLayout(pages, layouts, "app", nil)
	if !hasDiag(diags, "[TM-13]") {
		t.Errorf("expected TM-13 for unused layout 'auth', got %v", diags)
	}
	count := countDiag(diags, "[TM-13]")
	if count != 1 {
		t.Errorf("expected 1 TM-13 diagnostic, got %d: %+v", count, diags)
	}
}
