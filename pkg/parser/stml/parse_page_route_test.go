//ff:func feature=stml-parse type=test control=sequence
//ff:what TestParsePageWithDataRoute — data-route 속성 추출 검증

package stml

import (
	"strings"
	"testing"
)

func TestParsePageWithDataRoute(t *testing.T) {
	input := `<main data-route="/buildings/:buildingId/units/:id">
  <article data-fetch="GetUnit" data-param-buildingId="route.buildingId" data-param-id="route.id">
    <h2 data-field="Name"></h2>
  </article>
</main>`

	page, diags := ParseReader("unit-detail.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if page.Route != "/buildings/:buildingId/units/:id" {
		t.Errorf("Route = %q, want %q", page.Route, "/buildings/:buildingId/units/:id")
	}
}
