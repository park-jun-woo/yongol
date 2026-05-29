//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM12_EmptyDefaultLayout_Skip

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM12_EmptyDefaultLayout_Skip(t *testing.T) {
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm12DefaultLayoutNotFound("", layouts)
	if hasDiag(diags, "[TM-12]") {
		t.Errorf("unexpected TM-12 for empty defaultLayout, got %v", diags)
	}
}

// ---------- TM-13 ----------
