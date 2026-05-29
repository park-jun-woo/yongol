//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParsePageWithBothLayoutAndRoute — data-layout + data-route 동시 추출 검증

package stml

import (
	"strings"
	"testing"
)

func TestParsePageWithBothLayoutAndRoute(t *testing.T) {
	input := `<main data-layout="app" data-route="/buildings/:buildingId/units/:id">
  <article data-fetch="GetUnit" data-param-buildingId="route.buildingId" data-param-id="route.id">
    <h2 data-field="Name"></h2>
  </article>
</main>`

	page, diags := ParseReader("unit-detail.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if page.Layout != "app" {
		t.Errorf("Layout = %q, want %q", page.Layout, "app")
	}
	if page.Route != "/buildings/:buildingId/units/:id" {
		t.Errorf("Route = %q, want %q", page.Route, "/buildings/:buildingId/units/:id")
	}
}
