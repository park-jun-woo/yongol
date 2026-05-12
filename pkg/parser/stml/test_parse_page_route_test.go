//ff:func feature=stml-parse type=test control=sequence
//ff:what 페이지 data-route 속성 추출 테스트

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

func TestParsePageWithoutDataRoute(t *testing.T) {
	input := `<main>
  <div data-action="Login">
    <input data-field="Email" type="email" />
    <button type="submit">Login</button>
  </div>
</main>`

	page, diags := ParseReader("login.html", strings.NewReader(input))
	if len(diags) > 0 {
		t.Fatal(diags)
	}
	if page.Route != "" {
		t.Errorf("Route = %q, want empty string", page.Route)
	}
}

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
