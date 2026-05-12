//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM11_MultiplePages_CountsAll

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM11_MultiplePages_CountsAll(t *testing.T) {
	pages := []stml.PageSpec{
		{Name: "page-a", FileName: "page-a.html", Layout: "missing1"},
		{Name: "page-b", FileName: "page-b.html", Layout: "missing2"},
		{Name: "page-c", FileName: "page-c.html", Layout: "app"},
	}
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm11LayoutNotFound(pages, layouts)
	count := countDiag(diags, "[TM-11]")
	if count != 2 {
		t.Errorf("expected 2 TM-11 diagnostics, got %d: %+v", count, diags)
	}
}

// ---------- TM-12 ----------
