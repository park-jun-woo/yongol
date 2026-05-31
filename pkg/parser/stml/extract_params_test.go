//ff:func feature=stml-parse type=test control=iteration dimension=1
//ff:what hasDescendantData/hasDescendantField/hasDescendantDataInFetch/extractParams/extractPageMetaAttrs
package stml

import (
	"testing"
)

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
