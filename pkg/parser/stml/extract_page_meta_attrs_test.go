//ff:func feature=stml-parse type=test control=sequence
//ff:what hasDescendantData/hasDescendantField/hasDescendantDataInFetch/extractParams/extractPageMetaAttrs
package stml

import (
	"testing"
)

func TestExtractPageMetaAttrs(t *testing.T) {
	n := firstElementNode(t, `<div data-layout="main" data-route="/users"></div>`, "div")
	page := &PageSpec{}
	extractPageMetaAttrs(n, page)
	if page.Layout != "main" || page.Route != "/users" {
		t.Errorf("page = %+v", page)
	}
	// existing values are not overwritten
	page2 := &PageSpec{Layout: "kept", Route: "/kept"}
	extractPageMetaAttrs(n, page2)
	if page2.Layout != "kept" || page2.Route != "/kept" {
		t.Errorf("should not overwrite: %+v", page2)
	}
}
