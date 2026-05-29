//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM12_DefaultLayoutNotFound_Error

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM12_DefaultLayoutNotFound_Error(t *testing.T) {
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm12DefaultLayoutNotFound("main", layouts)
	if !hasDiag(diags, "[TM-12]") {
		t.Errorf("expected TM-12 for missing defaultLayout 'main', got %v", diags)
	}
}
