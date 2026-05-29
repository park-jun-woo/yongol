//ff:func feature=validate type=test control=sequence topic=stml-openapi
//ff:what TestTM12_DefaultLayoutFound_Pass

package stml_openapi

import (
	"testing"

	"github.com/park-jun-woo/yongol/pkg/parser/stml"
)

func TestTM12_DefaultLayoutFound_Pass(t *testing.T) {
	layouts := []stml.LayoutSpec{{Name: "app", File: "layouts/app.html"}}
	diags := tm12DefaultLayoutNotFound("app", layouts)
	if hasDiag(diags, "[TM-12]") {
		t.Errorf("unexpected TM-12 for valid defaultLayout, got %v", diags)
	}
}
