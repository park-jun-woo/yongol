//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM13_NoPages_AllUnused

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM13_NoPages_AllUnused(t *testing.T) {
	layouts := []stml.LayoutSpec{
		{Name: "app", File: "layouts/app.html"},
		{Name: "auth", File: "layouts/auth.html"},
	}
	diags := tm13UnusedLayout(nil, layouts, "", nil)
	count := countDiag(diags, "[TM-13]")
	if count != 2 {
		t.Errorf("expected 2 TM-13 diagnostics, got %d: %+v", count, diags)
	}
}

// ---------- run.go integration: skip when no layouts ----------
