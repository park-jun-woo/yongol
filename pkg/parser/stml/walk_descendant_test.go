//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what hasDescendantData/hasDescendantField/hasDescendantDataInFetch/extractParams/extractPageMetaAttrs

package stml

import "testing"

func TestHasDescendantData(t *testing.T) {
	with := firstElementNode(t, `<div><section><span data-fetch="/a"></span></section></div>`, "div")
	if !hasDescendantData(with) {
		t.Errorf("expected descendant data-fetch")
	}
	without := firstElementNode(t, `<div><section><span class="c"></span></section></div>`, "div")
	if hasDescendantData(without) {
		t.Errorf("expected no descendant data")
	}
}

func TestHasDescendantField(t *testing.T) {
	with := firstElementNode(t, `<div><p><span data-field="name"></span></p></div>`, "div")
	if !hasDescendantField(with) {
		t.Errorf("expected descendant data-field")
	}
	without := firstElementNode(t, `<div><p>text</p></div>`, "div")
	if hasDescendantField(without) {
		t.Errorf("expected no descendant field")
	}
}

func TestHasDescendantDataInFetch(t *testing.T) {
	with := firstElementNode(t, `<div><ul><li data-field="x"></li></ul></div>`, "div")
	if !hasDescendantDataInFetch(with) {
		t.Errorf("expected descendant data-* attr")
	}
	without := firstElementNode(t, `<div><ul><li class="c"></li></ul></div>`, "div")
	if hasDescendantDataInFetch(without) {
		t.Errorf("expected no descendant data-* attr")
	}
}

func TestExtractParams(t *testing.T) {
	n := firstElementNode(t, `<div data-param-reservation-id="route.ReservationID" data-param-page="query.Page" class="c"></div>`, "div")
	params := extractParams(n)
	if len(params) != 2 {
		t.Fatalf("len = %d, want 2 (%+v)", len(params), params)
	}
	got := map[string]string{}
	for _, p := range params {
		got[p.Name] = p.Source
	}
	if got["reservationId"] != "route.ReservationID" {
		t.Errorf("reservationId source = %q (got %+v)", got["reservationId"], got)
	}
	if got["page"] != "query.Page" {
		t.Errorf("page source = %q", got["page"])
	}
	// no params
	none := firstElementNode(t, `<div class="c"></div>`, "div")
	if p := extractParams(none); p != nil {
		t.Errorf("expected nil, got %+v", p)
	}
}

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
